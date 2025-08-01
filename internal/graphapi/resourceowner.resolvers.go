package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.78

import (
	"context"

	"entgo.io/contrib/entgql"
	"go.infratographer.com/metadata-api/internal/ent/generated"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotationnamespace"
	"go.infratographer.com/metadata-api/internal/ent/generated/metadata"
	"go.infratographer.com/x/gidx"
)

// Owner is the resolver for the owner field.
func (r *annotationNamespaceResolver) Owner(ctx context.Context, obj *generated.AnnotationNamespace) (*ResourceOwner, error) {
	return &ResourceOwner{ID: obj.ID}, nil
}

// AnnotationNamespaces is the resolver for the annotationNamespaces field.
func (r *resourceOwnerResolver) AnnotationNamespaces(ctx context.Context, obj *ResourceOwner, after *entgql.Cursor[gidx.PrefixedID], first *int, before *entgql.Cursor[gidx.PrefixedID], last *int, orderBy *generated.AnnotationNamespaceOrder, where *generated.AnnotationNamespaceWhereInput) (*generated.AnnotationNamespaceConnection, error) {
	return r.client.AnnotationNamespace.Query().Where(annotationnamespace.OwnerID(obj.ID)).Paginate(ctx, after, first, before, last, generated.WithAnnotationNamespaceOrder(orderBy), generated.WithAnnotationNamespaceFilter(where.Filter))
}

// Metadata is the resolver for the metadata field.
func (r *resourceOwnerResolver) Metadata(ctx context.Context, obj *ResourceOwner) (*generated.Metadata, error) {
	m, err := r.client.Metadata.Query().Where(metadata.NodeID(obj.ID)).First(ctx)
	// Don't return an error if it isn't found, metadata is optional
	err = generated.MaskNotFound(err)
	return m, err
}

// ResourceOwner returns ResourceOwnerResolver implementation.
func (r *Resolver) ResourceOwner() ResourceOwnerResolver { return &resourceOwnerResolver{r} }

type resourceOwnerResolver struct{ *Resolver }
