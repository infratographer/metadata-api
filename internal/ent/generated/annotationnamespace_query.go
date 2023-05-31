// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotation"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotationnamespace"
	"go.infratographer.com/metadata-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// AnnotationNamespaceQuery is the builder for querying AnnotationNamespace entities.
type AnnotationNamespaceQuery struct {
	config
	ctx                  *QueryContext
	order                []annotationnamespace.OrderOption
	inters               []Interceptor
	predicates           []predicate.AnnotationNamespace
	withAnnotations      *AnnotationQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*AnnotationNamespace) error
	withNamedAnnotations map[string]*AnnotationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AnnotationNamespaceQuery builder.
func (anq *AnnotationNamespaceQuery) Where(ps ...predicate.AnnotationNamespace) *AnnotationNamespaceQuery {
	anq.predicates = append(anq.predicates, ps...)
	return anq
}

// Limit the number of records to be returned by this query.
func (anq *AnnotationNamespaceQuery) Limit(limit int) *AnnotationNamespaceQuery {
	anq.ctx.Limit = &limit
	return anq
}

// Offset to start from.
func (anq *AnnotationNamespaceQuery) Offset(offset int) *AnnotationNamespaceQuery {
	anq.ctx.Offset = &offset
	return anq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (anq *AnnotationNamespaceQuery) Unique(unique bool) *AnnotationNamespaceQuery {
	anq.ctx.Unique = &unique
	return anq
}

// Order specifies how the records should be ordered.
func (anq *AnnotationNamespaceQuery) Order(o ...annotationnamespace.OrderOption) *AnnotationNamespaceQuery {
	anq.order = append(anq.order, o...)
	return anq
}

// QueryAnnotations chains the current query on the "annotations" edge.
func (anq *AnnotationNamespaceQuery) QueryAnnotations() *AnnotationQuery {
	query := (&AnnotationClient{config: anq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := anq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := anq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(annotationnamespace.Table, annotationnamespace.FieldID, selector),
			sqlgraph.To(annotation.Table, annotation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, annotationnamespace.AnnotationsTable, annotationnamespace.AnnotationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(anq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AnnotationNamespace entity from the query.
// Returns a *NotFoundError when no AnnotationNamespace was found.
func (anq *AnnotationNamespaceQuery) First(ctx context.Context) (*AnnotationNamespace, error) {
	nodes, err := anq.Limit(1).All(setContextOp(ctx, anq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{annotationnamespace.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (anq *AnnotationNamespaceQuery) FirstX(ctx context.Context) *AnnotationNamespace {
	node, err := anq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AnnotationNamespace ID from the query.
// Returns a *NotFoundError when no AnnotationNamespace ID was found.
func (anq *AnnotationNamespaceQuery) FirstID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = anq.Limit(1).IDs(setContextOp(ctx, anq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{annotationnamespace.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (anq *AnnotationNamespaceQuery) FirstIDX(ctx context.Context) gidx.PrefixedID {
	id, err := anq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AnnotationNamespace entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AnnotationNamespace entity is found.
// Returns a *NotFoundError when no AnnotationNamespace entities are found.
func (anq *AnnotationNamespaceQuery) Only(ctx context.Context) (*AnnotationNamespace, error) {
	nodes, err := anq.Limit(2).All(setContextOp(ctx, anq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{annotationnamespace.Label}
	default:
		return nil, &NotSingularError{annotationnamespace.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (anq *AnnotationNamespaceQuery) OnlyX(ctx context.Context) *AnnotationNamespace {
	node, err := anq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AnnotationNamespace ID in the query.
// Returns a *NotSingularError when more than one AnnotationNamespace ID is found.
// Returns a *NotFoundError when no entities are found.
func (anq *AnnotationNamespaceQuery) OnlyID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = anq.Limit(2).IDs(setContextOp(ctx, anq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{annotationnamespace.Label}
	default:
		err = &NotSingularError{annotationnamespace.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (anq *AnnotationNamespaceQuery) OnlyIDX(ctx context.Context) gidx.PrefixedID {
	id, err := anq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AnnotationNamespaces.
func (anq *AnnotationNamespaceQuery) All(ctx context.Context) ([]*AnnotationNamespace, error) {
	ctx = setContextOp(ctx, anq.ctx, "All")
	if err := anq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AnnotationNamespace, *AnnotationNamespaceQuery]()
	return withInterceptors[[]*AnnotationNamespace](ctx, anq, qr, anq.inters)
}

// AllX is like All, but panics if an error occurs.
func (anq *AnnotationNamespaceQuery) AllX(ctx context.Context) []*AnnotationNamespace {
	nodes, err := anq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AnnotationNamespace IDs.
func (anq *AnnotationNamespaceQuery) IDs(ctx context.Context) (ids []gidx.PrefixedID, err error) {
	if anq.ctx.Unique == nil && anq.path != nil {
		anq.Unique(true)
	}
	ctx = setContextOp(ctx, anq.ctx, "IDs")
	if err = anq.Select(annotationnamespace.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (anq *AnnotationNamespaceQuery) IDsX(ctx context.Context) []gidx.PrefixedID {
	ids, err := anq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (anq *AnnotationNamespaceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, anq.ctx, "Count")
	if err := anq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, anq, querierCount[*AnnotationNamespaceQuery](), anq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (anq *AnnotationNamespaceQuery) CountX(ctx context.Context) int {
	count, err := anq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (anq *AnnotationNamespaceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, anq.ctx, "Exist")
	switch _, err := anq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (anq *AnnotationNamespaceQuery) ExistX(ctx context.Context) bool {
	exist, err := anq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AnnotationNamespaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (anq *AnnotationNamespaceQuery) Clone() *AnnotationNamespaceQuery {
	if anq == nil {
		return nil
	}
	return &AnnotationNamespaceQuery{
		config:          anq.config,
		ctx:             anq.ctx.Clone(),
		order:           append([]annotationnamespace.OrderOption{}, anq.order...),
		inters:          append([]Interceptor{}, anq.inters...),
		predicates:      append([]predicate.AnnotationNamespace{}, anq.predicates...),
		withAnnotations: anq.withAnnotations.Clone(),
		// clone intermediate query.
		sql:  anq.sql.Clone(),
		path: anq.path,
	}
}

// WithAnnotations tells the query-builder to eager-load the nodes that are connected to
// the "annotations" edge. The optional arguments are used to configure the query builder of the edge.
func (anq *AnnotationNamespaceQuery) WithAnnotations(opts ...func(*AnnotationQuery)) *AnnotationNamespaceQuery {
	query := (&AnnotationClient{config: anq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	anq.withAnnotations = query
	return anq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AnnotationNamespace.Query().
//		GroupBy(annotationnamespace.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (anq *AnnotationNamespaceQuery) GroupBy(field string, fields ...string) *AnnotationNamespaceGroupBy {
	anq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AnnotationNamespaceGroupBy{build: anq}
	grbuild.flds = &anq.ctx.Fields
	grbuild.label = annotationnamespace.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.AnnotationNamespace.Query().
//		Select(annotationnamespace.FieldCreatedAt).
//		Scan(ctx, &v)
func (anq *AnnotationNamespaceQuery) Select(fields ...string) *AnnotationNamespaceSelect {
	anq.ctx.Fields = append(anq.ctx.Fields, fields...)
	sbuild := &AnnotationNamespaceSelect{AnnotationNamespaceQuery: anq}
	sbuild.label = annotationnamespace.Label
	sbuild.flds, sbuild.scan = &anq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AnnotationNamespaceSelect configured with the given aggregations.
func (anq *AnnotationNamespaceQuery) Aggregate(fns ...AggregateFunc) *AnnotationNamespaceSelect {
	return anq.Select().Aggregate(fns...)
}

func (anq *AnnotationNamespaceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range anq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, anq); err != nil {
				return err
			}
		}
	}
	for _, f := range anq.ctx.Fields {
		if !annotationnamespace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if anq.path != nil {
		prev, err := anq.path(ctx)
		if err != nil {
			return err
		}
		anq.sql = prev
	}
	return nil
}

func (anq *AnnotationNamespaceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AnnotationNamespace, error) {
	var (
		nodes       = []*AnnotationNamespace{}
		_spec       = anq.querySpec()
		loadedTypes = [1]bool{
			anq.withAnnotations != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AnnotationNamespace).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AnnotationNamespace{config: anq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(anq.modifiers) > 0 {
		_spec.Modifiers = anq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, anq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := anq.withAnnotations; query != nil {
		if err := anq.loadAnnotations(ctx, query, nodes,
			func(n *AnnotationNamespace) { n.Edges.Annotations = []*Annotation{} },
			func(n *AnnotationNamespace, e *Annotation) { n.Edges.Annotations = append(n.Edges.Annotations, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range anq.withNamedAnnotations {
		if err := anq.loadAnnotations(ctx, query, nodes,
			func(n *AnnotationNamespace) { n.appendNamedAnnotations(name) },
			func(n *AnnotationNamespace, e *Annotation) { n.appendNamedAnnotations(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range anq.loadTotal {
		if err := anq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (anq *AnnotationNamespaceQuery) loadAnnotations(ctx context.Context, query *AnnotationQuery, nodes []*AnnotationNamespace, init func(*AnnotationNamespace), assign func(*AnnotationNamespace, *Annotation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID]*AnnotationNamespace)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(annotation.FieldAnnotationNamespaceID)
	}
	query.Where(predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(annotationnamespace.AnnotationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AnnotationNamespaceID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "annotation_namespace_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (anq *AnnotationNamespaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := anq.querySpec()
	if len(anq.modifiers) > 0 {
		_spec.Modifiers = anq.modifiers
	}
	_spec.Node.Columns = anq.ctx.Fields
	if len(anq.ctx.Fields) > 0 {
		_spec.Unique = anq.ctx.Unique != nil && *anq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, anq.driver, _spec)
}

func (anq *AnnotationNamespaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(annotationnamespace.Table, annotationnamespace.Columns, sqlgraph.NewFieldSpec(annotationnamespace.FieldID, field.TypeString))
	_spec.From = anq.sql
	if unique := anq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if anq.path != nil {
		_spec.Unique = true
	}
	if fields := anq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, annotationnamespace.FieldID)
		for i := range fields {
			if fields[i] != annotationnamespace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := anq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := anq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := anq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := anq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (anq *AnnotationNamespaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(anq.driver.Dialect())
	t1 := builder.Table(annotationnamespace.Table)
	columns := anq.ctx.Fields
	if len(columns) == 0 {
		columns = annotationnamespace.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if anq.sql != nil {
		selector = anq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if anq.ctx.Unique != nil && *anq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range anq.predicates {
		p(selector)
	}
	for _, p := range anq.order {
		p(selector)
	}
	if offset := anq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := anq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAnnotations tells the query-builder to eager-load the nodes that are connected to the "annotations"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (anq *AnnotationNamespaceQuery) WithNamedAnnotations(name string, opts ...func(*AnnotationQuery)) *AnnotationNamespaceQuery {
	query := (&AnnotationClient{config: anq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if anq.withNamedAnnotations == nil {
		anq.withNamedAnnotations = make(map[string]*AnnotationQuery)
	}
	anq.withNamedAnnotations[name] = query
	return anq
}

// AnnotationNamespaceGroupBy is the group-by builder for AnnotationNamespace entities.
type AnnotationNamespaceGroupBy struct {
	selector
	build *AnnotationNamespaceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (angb *AnnotationNamespaceGroupBy) Aggregate(fns ...AggregateFunc) *AnnotationNamespaceGroupBy {
	angb.fns = append(angb.fns, fns...)
	return angb
}

// Scan applies the selector query and scans the result into the given value.
func (angb *AnnotationNamespaceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, angb.build.ctx, "GroupBy")
	if err := angb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AnnotationNamespaceQuery, *AnnotationNamespaceGroupBy](ctx, angb.build, angb, angb.build.inters, v)
}

func (angb *AnnotationNamespaceGroupBy) sqlScan(ctx context.Context, root *AnnotationNamespaceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(angb.fns))
	for _, fn := range angb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*angb.flds)+len(angb.fns))
		for _, f := range *angb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*angb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := angb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AnnotationNamespaceSelect is the builder for selecting fields of AnnotationNamespace entities.
type AnnotationNamespaceSelect struct {
	*AnnotationNamespaceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ans *AnnotationNamespaceSelect) Aggregate(fns ...AggregateFunc) *AnnotationNamespaceSelect {
	ans.fns = append(ans.fns, fns...)
	return ans
}

// Scan applies the selector query and scans the result into the given value.
func (ans *AnnotationNamespaceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ans.ctx, "Select")
	if err := ans.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AnnotationNamespaceQuery, *AnnotationNamespaceSelect](ctx, ans.AnnotationNamespaceQuery, ans, ans.inters, v)
}

func (ans *AnnotationNamespaceSelect) sqlScan(ctx context.Context, root *AnnotationNamespaceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ans.fns))
	for _, fn := range ans.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ans.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ans.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
