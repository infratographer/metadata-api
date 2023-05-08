// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/metadata-api/internal/ent/generated"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotation"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotationnamespace"
	"go.infratographer.com/metadata-api/internal/ent/generated/metadata"
	"go.infratographer.com/metadata-api/internal/ent/generated/predicate"
	"go.infratographer.com/metadata-api/internal/ent/generated/status"
	"go.infratographer.com/metadata-api/internal/ent/generated/statusnamespace"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next generated.Querier) generated.Querier {
	return generated.QuerierFunc(func(ctx context.Context, q generated.Query) (generated.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q generated.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The AnnotationFunc type is an adapter to allow the use of ordinary function as a Querier.
type AnnotationFunc func(context.Context, *generated.AnnotationQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f AnnotationFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.AnnotationQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.AnnotationQuery", q)
}

// The TraverseAnnotation type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAnnotation func(context.Context, *generated.AnnotationQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAnnotation) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAnnotation) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.AnnotationQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.AnnotationQuery", q)
}

// The AnnotationNamespaceFunc type is an adapter to allow the use of ordinary function as a Querier.
type AnnotationNamespaceFunc func(context.Context, *generated.AnnotationNamespaceQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f AnnotationNamespaceFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.AnnotationNamespaceQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.AnnotationNamespaceQuery", q)
}

// The TraverseAnnotationNamespace type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAnnotationNamespace func(context.Context, *generated.AnnotationNamespaceQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAnnotationNamespace) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAnnotationNamespace) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.AnnotationNamespaceQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.AnnotationNamespaceQuery", q)
}

// The MetadataFunc type is an adapter to allow the use of ordinary function as a Querier.
type MetadataFunc func(context.Context, *generated.MetadataQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f MetadataFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.MetadataQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.MetadataQuery", q)
}

// The TraverseMetadata type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMetadata func(context.Context, *generated.MetadataQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMetadata) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMetadata) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.MetadataQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.MetadataQuery", q)
}

// The StatusFunc type is an adapter to allow the use of ordinary function as a Querier.
type StatusFunc func(context.Context, *generated.StatusQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f StatusFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.StatusQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.StatusQuery", q)
}

// The TraverseStatus type is an adapter to allow the use of ordinary function as Traverser.
type TraverseStatus func(context.Context, *generated.StatusQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseStatus) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseStatus) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.StatusQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.StatusQuery", q)
}

// The StatusNamespaceFunc type is an adapter to allow the use of ordinary function as a Querier.
type StatusNamespaceFunc func(context.Context, *generated.StatusNamespaceQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f StatusNamespaceFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.StatusNamespaceQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.StatusNamespaceQuery", q)
}

// The TraverseStatusNamespace type is an adapter to allow the use of ordinary function as Traverser.
type TraverseStatusNamespace func(context.Context, *generated.StatusNamespaceQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseStatusNamespace) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseStatusNamespace) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.StatusNamespaceQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.StatusNamespaceQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q generated.Query) (Query, error) {
	switch q := q.(type) {
	case *generated.AnnotationQuery:
		return &query[*generated.AnnotationQuery, predicate.Annotation, annotation.OrderOption]{typ: generated.TypeAnnotation, tq: q}, nil
	case *generated.AnnotationNamespaceQuery:
		return &query[*generated.AnnotationNamespaceQuery, predicate.AnnotationNamespace, annotationnamespace.OrderOption]{typ: generated.TypeAnnotationNamespace, tq: q}, nil
	case *generated.MetadataQuery:
		return &query[*generated.MetadataQuery, predicate.Metadata, metadata.OrderOption]{typ: generated.TypeMetadata, tq: q}, nil
	case *generated.StatusQuery:
		return &query[*generated.StatusQuery, predicate.Status, status.OrderOption]{typ: generated.TypeStatus, tq: q}, nil
	case *generated.StatusNamespaceQuery:
		return &query[*generated.StatusNamespaceQuery, predicate.StatusNamespace, statusnamespace.OrderOption]{typ: generated.TypeStatusNamespace, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
