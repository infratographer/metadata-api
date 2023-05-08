package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"

	"entgo.io/contrib/entgql"
	"go.infratographer.com/metadata-api/internal/ent/generated"
	"go.infratographer.com/metadata-api/internal/ent/generated/statusnamespace"
	"go.infratographer.com/x/gidx"
)

// StatusNamespaces is the resolver for the statusNamespaces field.
func (r *resourceProviderResolver) StatusNamespaces(ctx context.Context, obj *ResourceProvider, after *entgql.Cursor[gidx.PrefixedID], first *int, before *entgql.Cursor[gidx.PrefixedID], last *int, orderBy *generated.StatusNamespaceOrder, where *generated.StatusNamespaceWhereInput) (*generated.StatusNamespaceConnection, error) {
	return r.client.StatusNamespace.Query().Where(statusnamespace.ResourceProviderID(obj.ID)).Paginate(ctx, after, first, before, last, generated.WithStatusNamespaceOrder(orderBy), generated.WithStatusNamespaceFilter(where.Filter))
}

// ResourceProvider is the resolver for the resourceProvider field.
func (r *statusNamespaceResolver) ResourceProvider(ctx context.Context, obj *generated.StatusNamespace) (*ResourceProvider, error) {
	return &ResourceProvider{ID: obj.ResourceProviderID}, nil
}

// ResourceProvider returns ResourceProviderResolver implementation.
func (r *Resolver) ResourceProvider() ResourceProviderResolver { return &resourceProviderResolver{r} }

type resourceProviderResolver struct{ *Resolver }
