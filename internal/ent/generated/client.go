// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"go.infratographer.com/metadata-api/internal/ent/generated/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotation"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotationnamespace"
	"go.infratographer.com/metadata-api/internal/ent/generated/metadata"
	"go.infratographer.com/metadata-api/internal/ent/generated/status"
	"go.infratographer.com/metadata-api/internal/ent/generated/statusnamespace"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/gidx"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Annotation is the client for interacting with the Annotation builders.
	Annotation *AnnotationClient
	// AnnotationNamespace is the client for interacting with the AnnotationNamespace builders.
	AnnotationNamespace *AnnotationNamespaceClient
	// Metadata is the client for interacting with the Metadata builders.
	Metadata *MetadataClient
	// Status is the client for interacting with the Status builders.
	Status *StatusClient
	// StatusNamespace is the client for interacting with the StatusNamespace builders.
	StatusNamespace *StatusNamespaceClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Annotation = NewAnnotationClient(c.config)
	c.AnnotationNamespace = NewAnnotationNamespaceClient(c.config)
	c.Metadata = NewMetadataClient(c.config)
	c.Status = NewStatusClient(c.config)
	c.StatusNamespace = NewStatusNamespaceClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters          *inters
		EventsPublisher events.Connection
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// EventsPublisher configures the EventsPublisher.
func EventsPublisher(v events.Connection) Option {
	return func(c *config) {
		c.EventsPublisher = v
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("generated: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("generated: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		Annotation:          NewAnnotationClient(cfg),
		AnnotationNamespace: NewAnnotationNamespaceClient(cfg),
		Metadata:            NewMetadataClient(cfg),
		Status:              NewStatusClient(cfg),
		StatusNamespace:     NewStatusNamespaceClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		Annotation:          NewAnnotationClient(cfg),
		AnnotationNamespace: NewAnnotationNamespaceClient(cfg),
		Metadata:            NewMetadataClient(cfg),
		Status:              NewStatusClient(cfg),
		StatusNamespace:     NewStatusNamespaceClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Annotation.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Annotation.Use(hooks...)
	c.AnnotationNamespace.Use(hooks...)
	c.Metadata.Use(hooks...)
	c.Status.Use(hooks...)
	c.StatusNamespace.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Annotation.Intercept(interceptors...)
	c.AnnotationNamespace.Intercept(interceptors...)
	c.Metadata.Intercept(interceptors...)
	c.Status.Intercept(interceptors...)
	c.StatusNamespace.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AnnotationMutation:
		return c.Annotation.mutate(ctx, m)
	case *AnnotationNamespaceMutation:
		return c.AnnotationNamespace.mutate(ctx, m)
	case *MetadataMutation:
		return c.Metadata.mutate(ctx, m)
	case *StatusMutation:
		return c.Status.mutate(ctx, m)
	case *StatusNamespaceMutation:
		return c.StatusNamespace.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("generated: unknown mutation type %T", m)
	}
}

// AnnotationClient is a client for the Annotation schema.
type AnnotationClient struct {
	config
}

// NewAnnotationClient returns a client for the Annotation from the given config.
func NewAnnotationClient(c config) *AnnotationClient {
	return &AnnotationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `annotation.Hooks(f(g(h())))`.
func (c *AnnotationClient) Use(hooks ...Hook) {
	c.hooks.Annotation = append(c.hooks.Annotation, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `annotation.Intercept(f(g(h())))`.
func (c *AnnotationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Annotation = append(c.inters.Annotation, interceptors...)
}

// Create returns a builder for creating a Annotation entity.
func (c *AnnotationClient) Create() *AnnotationCreate {
	mutation := newAnnotationMutation(c.config, OpCreate)
	return &AnnotationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Annotation entities.
func (c *AnnotationClient) CreateBulk(builders ...*AnnotationCreate) *AnnotationCreateBulk {
	return &AnnotationCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AnnotationClient) MapCreateBulk(slice any, setFunc func(*AnnotationCreate, int)) *AnnotationCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AnnotationCreateBulk{err: fmt.Errorf("calling to AnnotationClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AnnotationCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AnnotationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Annotation.
func (c *AnnotationClient) Update() *AnnotationUpdate {
	mutation := newAnnotationMutation(c.config, OpUpdate)
	return &AnnotationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AnnotationClient) UpdateOne(a *Annotation) *AnnotationUpdateOne {
	mutation := newAnnotationMutation(c.config, OpUpdateOne, withAnnotation(a))
	return &AnnotationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AnnotationClient) UpdateOneID(id gidx.PrefixedID) *AnnotationUpdateOne {
	mutation := newAnnotationMutation(c.config, OpUpdateOne, withAnnotationID(id))
	return &AnnotationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Annotation.
func (c *AnnotationClient) Delete() *AnnotationDelete {
	mutation := newAnnotationMutation(c.config, OpDelete)
	return &AnnotationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AnnotationClient) DeleteOne(a *Annotation) *AnnotationDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AnnotationClient) DeleteOneID(id gidx.PrefixedID) *AnnotationDeleteOne {
	builder := c.Delete().Where(annotation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AnnotationDeleteOne{builder}
}

// Query returns a query builder for Annotation.
func (c *AnnotationClient) Query() *AnnotationQuery {
	return &AnnotationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAnnotation},
		inters: c.Interceptors(),
	}
}

// Get returns a Annotation entity by its id.
func (c *AnnotationClient) Get(ctx context.Context, id gidx.PrefixedID) (*Annotation, error) {
	return c.Query().Where(annotation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AnnotationClient) GetX(ctx context.Context, id gidx.PrefixedID) *Annotation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryNamespace queries the namespace edge of a Annotation.
func (c *AnnotationClient) QueryNamespace(a *Annotation) *AnnotationNamespaceQuery {
	query := (&AnnotationNamespaceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(annotation.Table, annotation.FieldID, id),
			sqlgraph.To(annotationnamespace.Table, annotationnamespace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, annotation.NamespaceTable, annotation.NamespaceColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMetadata queries the metadata edge of a Annotation.
func (c *AnnotationClient) QueryMetadata(a *Annotation) *MetadataQuery {
	query := (&MetadataClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(annotation.Table, annotation.FieldID, id),
			sqlgraph.To(metadata.Table, metadata.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, annotation.MetadataTable, annotation.MetadataColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AnnotationClient) Hooks() []Hook {
	return c.hooks.Annotation
}

// Interceptors returns the client interceptors.
func (c *AnnotationClient) Interceptors() []Interceptor {
	return c.inters.Annotation
}

func (c *AnnotationClient) mutate(ctx context.Context, m *AnnotationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AnnotationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AnnotationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AnnotationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AnnotationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Annotation mutation op: %q", m.Op())
	}
}

// AnnotationNamespaceClient is a client for the AnnotationNamespace schema.
type AnnotationNamespaceClient struct {
	config
}

// NewAnnotationNamespaceClient returns a client for the AnnotationNamespace from the given config.
func NewAnnotationNamespaceClient(c config) *AnnotationNamespaceClient {
	return &AnnotationNamespaceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `annotationnamespace.Hooks(f(g(h())))`.
func (c *AnnotationNamespaceClient) Use(hooks ...Hook) {
	c.hooks.AnnotationNamespace = append(c.hooks.AnnotationNamespace, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `annotationnamespace.Intercept(f(g(h())))`.
func (c *AnnotationNamespaceClient) Intercept(interceptors ...Interceptor) {
	c.inters.AnnotationNamespace = append(c.inters.AnnotationNamespace, interceptors...)
}

// Create returns a builder for creating a AnnotationNamespace entity.
func (c *AnnotationNamespaceClient) Create() *AnnotationNamespaceCreate {
	mutation := newAnnotationNamespaceMutation(c.config, OpCreate)
	return &AnnotationNamespaceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AnnotationNamespace entities.
func (c *AnnotationNamespaceClient) CreateBulk(builders ...*AnnotationNamespaceCreate) *AnnotationNamespaceCreateBulk {
	return &AnnotationNamespaceCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AnnotationNamespaceClient) MapCreateBulk(slice any, setFunc func(*AnnotationNamespaceCreate, int)) *AnnotationNamespaceCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AnnotationNamespaceCreateBulk{err: fmt.Errorf("calling to AnnotationNamespaceClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AnnotationNamespaceCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AnnotationNamespaceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AnnotationNamespace.
func (c *AnnotationNamespaceClient) Update() *AnnotationNamespaceUpdate {
	mutation := newAnnotationNamespaceMutation(c.config, OpUpdate)
	return &AnnotationNamespaceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AnnotationNamespaceClient) UpdateOne(an *AnnotationNamespace) *AnnotationNamespaceUpdateOne {
	mutation := newAnnotationNamespaceMutation(c.config, OpUpdateOne, withAnnotationNamespace(an))
	return &AnnotationNamespaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AnnotationNamespaceClient) UpdateOneID(id gidx.PrefixedID) *AnnotationNamespaceUpdateOne {
	mutation := newAnnotationNamespaceMutation(c.config, OpUpdateOne, withAnnotationNamespaceID(id))
	return &AnnotationNamespaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AnnotationNamespace.
func (c *AnnotationNamespaceClient) Delete() *AnnotationNamespaceDelete {
	mutation := newAnnotationNamespaceMutation(c.config, OpDelete)
	return &AnnotationNamespaceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AnnotationNamespaceClient) DeleteOne(an *AnnotationNamespace) *AnnotationNamespaceDeleteOne {
	return c.DeleteOneID(an.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AnnotationNamespaceClient) DeleteOneID(id gidx.PrefixedID) *AnnotationNamespaceDeleteOne {
	builder := c.Delete().Where(annotationnamespace.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AnnotationNamespaceDeleteOne{builder}
}

// Query returns a query builder for AnnotationNamespace.
func (c *AnnotationNamespaceClient) Query() *AnnotationNamespaceQuery {
	return &AnnotationNamespaceQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAnnotationNamespace},
		inters: c.Interceptors(),
	}
}

// Get returns a AnnotationNamespace entity by its id.
func (c *AnnotationNamespaceClient) Get(ctx context.Context, id gidx.PrefixedID) (*AnnotationNamespace, error) {
	return c.Query().Where(annotationnamespace.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AnnotationNamespaceClient) GetX(ctx context.Context, id gidx.PrefixedID) *AnnotationNamespace {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAnnotations queries the annotations edge of a AnnotationNamespace.
func (c *AnnotationNamespaceClient) QueryAnnotations(an *AnnotationNamespace) *AnnotationQuery {
	query := (&AnnotationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := an.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(annotationnamespace.Table, annotationnamespace.FieldID, id),
			sqlgraph.To(annotation.Table, annotation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, annotationnamespace.AnnotationsTable, annotationnamespace.AnnotationsColumn),
		)
		fromV = sqlgraph.Neighbors(an.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AnnotationNamespaceClient) Hooks() []Hook {
	return c.hooks.AnnotationNamespace
}

// Interceptors returns the client interceptors.
func (c *AnnotationNamespaceClient) Interceptors() []Interceptor {
	return c.inters.AnnotationNamespace
}

func (c *AnnotationNamespaceClient) mutate(ctx context.Context, m *AnnotationNamespaceMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AnnotationNamespaceCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AnnotationNamespaceUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AnnotationNamespaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AnnotationNamespaceDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown AnnotationNamespace mutation op: %q", m.Op())
	}
}

// MetadataClient is a client for the Metadata schema.
type MetadataClient struct {
	config
}

// NewMetadataClient returns a client for the Metadata from the given config.
func NewMetadataClient(c config) *MetadataClient {
	return &MetadataClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `metadata.Hooks(f(g(h())))`.
func (c *MetadataClient) Use(hooks ...Hook) {
	c.hooks.Metadata = append(c.hooks.Metadata, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `metadata.Intercept(f(g(h())))`.
func (c *MetadataClient) Intercept(interceptors ...Interceptor) {
	c.inters.Metadata = append(c.inters.Metadata, interceptors...)
}

// Create returns a builder for creating a Metadata entity.
func (c *MetadataClient) Create() *MetadataCreate {
	mutation := newMetadataMutation(c.config, OpCreate)
	return &MetadataCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Metadata entities.
func (c *MetadataClient) CreateBulk(builders ...*MetadataCreate) *MetadataCreateBulk {
	return &MetadataCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MetadataClient) MapCreateBulk(slice any, setFunc func(*MetadataCreate, int)) *MetadataCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MetadataCreateBulk{err: fmt.Errorf("calling to MetadataClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MetadataCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MetadataCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Metadata.
func (c *MetadataClient) Update() *MetadataUpdate {
	mutation := newMetadataMutation(c.config, OpUpdate)
	return &MetadataUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MetadataClient) UpdateOne(m *Metadata) *MetadataUpdateOne {
	mutation := newMetadataMutation(c.config, OpUpdateOne, withMetadata(m))
	return &MetadataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MetadataClient) UpdateOneID(id gidx.PrefixedID) *MetadataUpdateOne {
	mutation := newMetadataMutation(c.config, OpUpdateOne, withMetadataID(id))
	return &MetadataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Metadata.
func (c *MetadataClient) Delete() *MetadataDelete {
	mutation := newMetadataMutation(c.config, OpDelete)
	return &MetadataDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MetadataClient) DeleteOne(m *Metadata) *MetadataDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MetadataClient) DeleteOneID(id gidx.PrefixedID) *MetadataDeleteOne {
	builder := c.Delete().Where(metadata.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MetadataDeleteOne{builder}
}

// Query returns a query builder for Metadata.
func (c *MetadataClient) Query() *MetadataQuery {
	return &MetadataQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMetadata},
		inters: c.Interceptors(),
	}
}

// Get returns a Metadata entity by its id.
func (c *MetadataClient) Get(ctx context.Context, id gidx.PrefixedID) (*Metadata, error) {
	return c.Query().Where(metadata.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MetadataClient) GetX(ctx context.Context, id gidx.PrefixedID) *Metadata {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAnnotations queries the annotations edge of a Metadata.
func (c *MetadataClient) QueryAnnotations(m *Metadata) *AnnotationQuery {
	query := (&AnnotationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(metadata.Table, metadata.FieldID, id),
			sqlgraph.To(annotation.Table, annotation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, metadata.AnnotationsTable, metadata.AnnotationsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStatuses queries the statuses edge of a Metadata.
func (c *MetadataClient) QueryStatuses(m *Metadata) *StatusQuery {
	query := (&StatusClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(metadata.Table, metadata.FieldID, id),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, metadata.StatusesTable, metadata.StatusesColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MetadataClient) Hooks() []Hook {
	return c.hooks.Metadata
}

// Interceptors returns the client interceptors.
func (c *MetadataClient) Interceptors() []Interceptor {
	return c.inters.Metadata
}

func (c *MetadataClient) mutate(ctx context.Context, m *MetadataMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MetadataCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MetadataUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MetadataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MetadataDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Metadata mutation op: %q", m.Op())
	}
}

// StatusClient is a client for the Status schema.
type StatusClient struct {
	config
}

// NewStatusClient returns a client for the Status from the given config.
func NewStatusClient(c config) *StatusClient {
	return &StatusClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `status.Hooks(f(g(h())))`.
func (c *StatusClient) Use(hooks ...Hook) {
	c.hooks.Status = append(c.hooks.Status, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `status.Intercept(f(g(h())))`.
func (c *StatusClient) Intercept(interceptors ...Interceptor) {
	c.inters.Status = append(c.inters.Status, interceptors...)
}

// Create returns a builder for creating a Status entity.
func (c *StatusClient) Create() *StatusCreate {
	mutation := newStatusMutation(c.config, OpCreate)
	return &StatusCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Status entities.
func (c *StatusClient) CreateBulk(builders ...*StatusCreate) *StatusCreateBulk {
	return &StatusCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *StatusClient) MapCreateBulk(slice any, setFunc func(*StatusCreate, int)) *StatusCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &StatusCreateBulk{err: fmt.Errorf("calling to StatusClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*StatusCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &StatusCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Status.
func (c *StatusClient) Update() *StatusUpdate {
	mutation := newStatusMutation(c.config, OpUpdate)
	return &StatusUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StatusClient) UpdateOne(s *Status) *StatusUpdateOne {
	mutation := newStatusMutation(c.config, OpUpdateOne, withStatus(s))
	return &StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StatusClient) UpdateOneID(id gidx.PrefixedID) *StatusUpdateOne {
	mutation := newStatusMutation(c.config, OpUpdateOne, withStatusID(id))
	return &StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Status.
func (c *StatusClient) Delete() *StatusDelete {
	mutation := newStatusMutation(c.config, OpDelete)
	return &StatusDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StatusClient) DeleteOne(s *Status) *StatusDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StatusClient) DeleteOneID(id gidx.PrefixedID) *StatusDeleteOne {
	builder := c.Delete().Where(status.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StatusDeleteOne{builder}
}

// Query returns a query builder for Status.
func (c *StatusClient) Query() *StatusQuery {
	return &StatusQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStatus},
		inters: c.Interceptors(),
	}
}

// Get returns a Status entity by its id.
func (c *StatusClient) Get(ctx context.Context, id gidx.PrefixedID) (*Status, error) {
	return c.Query().Where(status.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StatusClient) GetX(ctx context.Context, id gidx.PrefixedID) *Status {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryNamespace queries the namespace edge of a Status.
func (c *StatusClient) QueryNamespace(s *Status) *StatusNamespaceQuery {
	query := (&StatusNamespaceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(status.Table, status.FieldID, id),
			sqlgraph.To(statusnamespace.Table, statusnamespace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, status.NamespaceTable, status.NamespaceColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMetadata queries the metadata edge of a Status.
func (c *StatusClient) QueryMetadata(s *Status) *MetadataQuery {
	query := (&MetadataClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(status.Table, status.FieldID, id),
			sqlgraph.To(metadata.Table, metadata.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, status.MetadataTable, status.MetadataColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StatusClient) Hooks() []Hook {
	return c.hooks.Status
}

// Interceptors returns the client interceptors.
func (c *StatusClient) Interceptors() []Interceptor {
	return c.inters.Status
}

func (c *StatusClient) mutate(ctx context.Context, m *StatusMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StatusCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StatusUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StatusDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Status mutation op: %q", m.Op())
	}
}

// StatusNamespaceClient is a client for the StatusNamespace schema.
type StatusNamespaceClient struct {
	config
}

// NewStatusNamespaceClient returns a client for the StatusNamespace from the given config.
func NewStatusNamespaceClient(c config) *StatusNamespaceClient {
	return &StatusNamespaceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `statusnamespace.Hooks(f(g(h())))`.
func (c *StatusNamespaceClient) Use(hooks ...Hook) {
	c.hooks.StatusNamespace = append(c.hooks.StatusNamespace, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `statusnamespace.Intercept(f(g(h())))`.
func (c *StatusNamespaceClient) Intercept(interceptors ...Interceptor) {
	c.inters.StatusNamespace = append(c.inters.StatusNamespace, interceptors...)
}

// Create returns a builder for creating a StatusNamespace entity.
func (c *StatusNamespaceClient) Create() *StatusNamespaceCreate {
	mutation := newStatusNamespaceMutation(c.config, OpCreate)
	return &StatusNamespaceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of StatusNamespace entities.
func (c *StatusNamespaceClient) CreateBulk(builders ...*StatusNamespaceCreate) *StatusNamespaceCreateBulk {
	return &StatusNamespaceCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *StatusNamespaceClient) MapCreateBulk(slice any, setFunc func(*StatusNamespaceCreate, int)) *StatusNamespaceCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &StatusNamespaceCreateBulk{err: fmt.Errorf("calling to StatusNamespaceClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*StatusNamespaceCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &StatusNamespaceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for StatusNamespace.
func (c *StatusNamespaceClient) Update() *StatusNamespaceUpdate {
	mutation := newStatusNamespaceMutation(c.config, OpUpdate)
	return &StatusNamespaceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StatusNamespaceClient) UpdateOne(sn *StatusNamespace) *StatusNamespaceUpdateOne {
	mutation := newStatusNamespaceMutation(c.config, OpUpdateOne, withStatusNamespace(sn))
	return &StatusNamespaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StatusNamespaceClient) UpdateOneID(id gidx.PrefixedID) *StatusNamespaceUpdateOne {
	mutation := newStatusNamespaceMutation(c.config, OpUpdateOne, withStatusNamespaceID(id))
	return &StatusNamespaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for StatusNamespace.
func (c *StatusNamespaceClient) Delete() *StatusNamespaceDelete {
	mutation := newStatusNamespaceMutation(c.config, OpDelete)
	return &StatusNamespaceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StatusNamespaceClient) DeleteOne(sn *StatusNamespace) *StatusNamespaceDeleteOne {
	return c.DeleteOneID(sn.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StatusNamespaceClient) DeleteOneID(id gidx.PrefixedID) *StatusNamespaceDeleteOne {
	builder := c.Delete().Where(statusnamespace.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StatusNamespaceDeleteOne{builder}
}

// Query returns a query builder for StatusNamespace.
func (c *StatusNamespaceClient) Query() *StatusNamespaceQuery {
	return &StatusNamespaceQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStatusNamespace},
		inters: c.Interceptors(),
	}
}

// Get returns a StatusNamespace entity by its id.
func (c *StatusNamespaceClient) Get(ctx context.Context, id gidx.PrefixedID) (*StatusNamespace, error) {
	return c.Query().Where(statusnamespace.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StatusNamespaceClient) GetX(ctx context.Context, id gidx.PrefixedID) *StatusNamespace {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStatuses queries the statuses edge of a StatusNamespace.
func (c *StatusNamespaceClient) QueryStatuses(sn *StatusNamespace) *StatusQuery {
	query := (&StatusClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := sn.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(statusnamespace.Table, statusnamespace.FieldID, id),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, statusnamespace.StatusesTable, statusnamespace.StatusesColumn),
		)
		fromV = sqlgraph.Neighbors(sn.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StatusNamespaceClient) Hooks() []Hook {
	return c.hooks.StatusNamespace
}

// Interceptors returns the client interceptors.
func (c *StatusNamespaceClient) Interceptors() []Interceptor {
	return c.inters.StatusNamespace
}

func (c *StatusNamespaceClient) mutate(ctx context.Context, m *StatusNamespaceMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StatusNamespaceCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StatusNamespaceUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StatusNamespaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StatusNamespaceDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown StatusNamespace mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Annotation, AnnotationNamespace, Metadata, Status, StatusNamespace []ent.Hook
	}
	inters struct {
		Annotation, AnnotationNamespace, Metadata, Status,
		StatusNamespace []ent.Interceptor
	}
)
