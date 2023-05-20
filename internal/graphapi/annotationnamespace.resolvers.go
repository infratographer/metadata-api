package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31-dev

import (
	"context"
	"fmt"

	"go.infratographer.com/metadata-api/internal/ent/generated"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotation"
	"go.infratographer.com/x/gidx"
)

// AnnotationNamespaceCreate is the resolver for the annotationNamespaceCreate field.
func (r *mutationResolver) AnnotationNamespaceCreate(ctx context.Context, input generated.CreateAnnotationNamespaceInput) (*AnnotationNamespaceCreatePayload, error) {
	// TODO: authz check here
	ns, err := r.client.AnnotationNamespace.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &AnnotationNamespaceCreatePayload{AnnotationNamespace: ns}, nil
}

// AnnotationNamespaceDelete is the resolver for the annotationNamespaceDelete field.
func (r *mutationResolver) AnnotationNamespaceDelete(ctx context.Context, id gidx.PrefixedID, force bool) (*AnnotationNamespaceDeletePayload, error) {
	// TODO: authz check here
	antCount, err := r.client.Annotation.Query().Where(annotation.AnnotationNamespaceID(id)).Count(ctx)
	if err != nil {
		return nil, err
	}

	if antCount != 0 {
		if force {
			antCount, err = r.client.Annotation.Delete().Where(annotation.AnnotationNamespaceID(id)).Exec(ctx)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("annotation namespace is in use and can't be deleted")
		}
	}

	err = r.client.AnnotationNamespace.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &AnnotationNamespaceDeletePayload{DeletedID: id, AnnotationDeletedCount: antCount}, nil
}

// AnnotationNamespaceUpdate is the resolver for the annotationNamespaceUpdate field.
func (r *mutationResolver) AnnotationNamespaceUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateAnnotationNamespaceInput) (*AnnotationNamespaceUpdatePayload, error) {
	// TODO: authz check here
	ns, err := r.client.AnnotationNamespace.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	ns, err = ns.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &AnnotationNamespaceUpdatePayload{AnnotationNamespace: ns}, nil
}

// AnnotationNamespace is the resolver for the annotationNamespace field.
func (r *queryResolver) AnnotationNamespace(ctx context.Context, id gidx.PrefixedID) (*generated.AnnotationNamespace, error) {
	// TODO: authz check here
	return r.client.AnnotationNamespace.Get(ctx, id)
}
