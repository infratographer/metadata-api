package cmd

import (
	"context"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/loggingx"
	"go.infratographer.com/x/otelx"
	"go.infratographer.com/x/versionx"
	"go.uber.org/zap"

	"go.infratographer.com/metadata-api/internal/config"
	ent "go.infratographer.com/metadata-api/internal/ent/generated"
	"go.infratographer.com/metadata-api/internal/graphapi"

	"go.infratographer.com/metadata-api/internal/ent/generated/eventhooks"
)

const (
	defaultAPIListenAddr = ":7905"
	shutdownTimeout      = 10 * time.Second
)

var (
	enablePlayground bool
	serveDevMode     bool
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Metadata Graph API",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	echox.MustViperFlags(viper.GetViper(), serveCmd.Flags(), defaultAPIListenAddr)
	echojwtx.MustViperFlags(viper.GetViper(), serveCmd.Flags())

	events.MustViperFlags(viper.GetViper(), serveCmd.Flags(), appName)
	permissions.MustViperFlags(viper.GetViper(), serveCmd.Flags())

	// only available as a CLI arg because it shouldn't be something that could accidentially end up in a config file or env var
	serveCmd.Flags().BoolVar(&serveDevMode, "dev", false, "dev mode: enables playground, disables all auth checks, sets CORS to allow all, pretty logging, etc.")
	serveCmd.Flags().BoolVar(&enablePlayground, "playground", false, "enable the graph playground")
}

func serve(ctx context.Context) error {
	if serveDevMode {
		enablePlayground = true
		config.AppConfig.Logging.Debug = true
		config.AppConfig.Logging.Pretty = true
		config.AppConfig.Server.WithMiddleware(middleware.CORS())
		// reinit the logger
		logger = loggingx.InitLogger(appName, config.AppConfig.Logging)
		// this is a hack, echojwt needs to be updated to go into AppConfig
		viper.Set("oidc.enabled", false)
	}

	events, err := events.NewConnection(config.AppConfig.Events, events.WithLogger(logger))
	if err != nil {
		logger.Fatalw("failed to initialize events", "error", err)
	}

	if err := otelx.InitTracer(config.AppConfig.Tracing, appName, logger); err != nil {
		logger.Fatalw("failed to initialize tracer", "error", err)
	}

	db, err := crdbx.NewDB(config.AppConfig.CRDB, config.AppConfig.Tracing.Enabled)
	if err != nil {
		logger.Fatalw("failed to connect to database", "error", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			logger.Errorw("failed to close database connection", "error", err)
		}
	}()

	entDB := entsql.OpenDB(dialect.Postgres, db)

	cOpts := []ent.Option{ent.Driver(entDB), ent.EventsPublisher(events)}

	if config.AppConfig.Logging.Debug {
		cOpts = append(cOpts,
			ent.Log(logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	client := ent.NewClient(cOpts...)
	defer func() {
		if err := client.Close(); err != nil {
			logger.Errorw("failed to close database client", "error", err)
		}
	}()

	eventhooks.EventHooks(client)

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		logger.Errorf("failed creating schema resources", zap.Error(err))
		return err
	}

	var middleware []echo.MiddlewareFunc

	// jwt auth middleware
	if viper.GetBool("oidc.enabled") {
		auth, err := echojwtx.NewAuth(ctx, config.AppConfig.OIDC)
		if err != nil {
			logger.Fatalw("failed to initialize jwt authentication", zap.Error(err))
		}

		middleware = append(middleware, auth.Middleware())
	}

	srv, err := echox.NewServer(logger.Desugar(), config.AppConfig.Server, versionx.BuildDetails(), echox.WithLoggingSkipper(echox.SkipDefaultEndpoints))
	if err != nil {
		logger.Error("failed to create server", zap.Error(err))
	}

	perms, err := permissions.New(config.AppConfig.Permissions,
		permissions.WithLogger(logger),
		permissions.WithDefaultChecker(permissions.DefaultAllowChecker),
		permissions.WithEventsPublisher(events),
	)
	if err != nil {
		logger.Fatal("failed to initialize permissions", zap.Error(err))
	}

	middleware = append(middleware, perms.Middleware())

	r := graphapi.NewResolver(client, logger.Named("resolvers"))
	handler := r.Handler(enablePlayground, middleware...)

	srv.AddHandler(handler)

	defer func() {
		ctx, cancel := context.WithTimeout(ctx, shutdownTimeout)
		defer cancel()

		_ = events.Shutdown(ctx)
	}()

	if err := srv.RunWithContext(ctx); err != nil {
		logger.Error("failed to run server", zap.Error(err))
	}

	return nil
}
