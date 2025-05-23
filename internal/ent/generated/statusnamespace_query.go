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
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/metadata-api/internal/ent/generated/predicate"
	"go.infratographer.com/metadata-api/internal/ent/generated/status"
	"go.infratographer.com/metadata-api/internal/ent/generated/statusnamespace"
	"go.infratographer.com/x/gidx"
)

// StatusNamespaceQuery is the builder for querying StatusNamespace entities.
type StatusNamespaceQuery struct {
	config
	ctx               *QueryContext
	order             []statusnamespace.OrderOption
	inters            []Interceptor
	predicates        []predicate.StatusNamespace
	withStatuses      *StatusQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*StatusNamespace) error
	withNamedStatuses map[string]*StatusQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StatusNamespaceQuery builder.
func (snq *StatusNamespaceQuery) Where(ps ...predicate.StatusNamespace) *StatusNamespaceQuery {
	snq.predicates = append(snq.predicates, ps...)
	return snq
}

// Limit the number of records to be returned by this query.
func (snq *StatusNamespaceQuery) Limit(limit int) *StatusNamespaceQuery {
	snq.ctx.Limit = &limit
	return snq
}

// Offset to start from.
func (snq *StatusNamespaceQuery) Offset(offset int) *StatusNamespaceQuery {
	snq.ctx.Offset = &offset
	return snq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (snq *StatusNamespaceQuery) Unique(unique bool) *StatusNamespaceQuery {
	snq.ctx.Unique = &unique
	return snq
}

// Order specifies how the records should be ordered.
func (snq *StatusNamespaceQuery) Order(o ...statusnamespace.OrderOption) *StatusNamespaceQuery {
	snq.order = append(snq.order, o...)
	return snq
}

// QueryStatuses chains the current query on the "statuses" edge.
func (snq *StatusNamespaceQuery) QueryStatuses() *StatusQuery {
	query := (&StatusClient{config: snq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := snq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := snq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statusnamespace.Table, statusnamespace.FieldID, selector),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, statusnamespace.StatusesTable, statusnamespace.StatusesColumn),
		)
		fromU = sqlgraph.SetNeighbors(snq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StatusNamespace entity from the query.
// Returns a *NotFoundError when no StatusNamespace was found.
func (snq *StatusNamespaceQuery) First(ctx context.Context) (*StatusNamespace, error) {
	nodes, err := snq.Limit(1).All(setContextOp(ctx, snq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{statusnamespace.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (snq *StatusNamespaceQuery) FirstX(ctx context.Context) *StatusNamespace {
	node, err := snq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StatusNamespace ID from the query.
// Returns a *NotFoundError when no StatusNamespace ID was found.
func (snq *StatusNamespaceQuery) FirstID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = snq.Limit(1).IDs(setContextOp(ctx, snq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{statusnamespace.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (snq *StatusNamespaceQuery) FirstIDX(ctx context.Context) gidx.PrefixedID {
	id, err := snq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StatusNamespace entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StatusNamespace entity is found.
// Returns a *NotFoundError when no StatusNamespace entities are found.
func (snq *StatusNamespaceQuery) Only(ctx context.Context) (*StatusNamespace, error) {
	nodes, err := snq.Limit(2).All(setContextOp(ctx, snq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{statusnamespace.Label}
	default:
		return nil, &NotSingularError{statusnamespace.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (snq *StatusNamespaceQuery) OnlyX(ctx context.Context) *StatusNamespace {
	node, err := snq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StatusNamespace ID in the query.
// Returns a *NotSingularError when more than one StatusNamespace ID is found.
// Returns a *NotFoundError when no entities are found.
func (snq *StatusNamespaceQuery) OnlyID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = snq.Limit(2).IDs(setContextOp(ctx, snq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{statusnamespace.Label}
	default:
		err = &NotSingularError{statusnamespace.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (snq *StatusNamespaceQuery) OnlyIDX(ctx context.Context) gidx.PrefixedID {
	id, err := snq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StatusNamespaces.
func (snq *StatusNamespaceQuery) All(ctx context.Context) ([]*StatusNamespace, error) {
	ctx = setContextOp(ctx, snq.ctx, ent.OpQueryAll)
	if err := snq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*StatusNamespace, *StatusNamespaceQuery]()
	return withInterceptors[[]*StatusNamespace](ctx, snq, qr, snq.inters)
}

// AllX is like All, but panics if an error occurs.
func (snq *StatusNamespaceQuery) AllX(ctx context.Context) []*StatusNamespace {
	nodes, err := snq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StatusNamespace IDs.
func (snq *StatusNamespaceQuery) IDs(ctx context.Context) (ids []gidx.PrefixedID, err error) {
	if snq.ctx.Unique == nil && snq.path != nil {
		snq.Unique(true)
	}
	ctx = setContextOp(ctx, snq.ctx, ent.OpQueryIDs)
	if err = snq.Select(statusnamespace.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (snq *StatusNamespaceQuery) IDsX(ctx context.Context) []gidx.PrefixedID {
	ids, err := snq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (snq *StatusNamespaceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, snq.ctx, ent.OpQueryCount)
	if err := snq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, snq, querierCount[*StatusNamespaceQuery](), snq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (snq *StatusNamespaceQuery) CountX(ctx context.Context) int {
	count, err := snq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (snq *StatusNamespaceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, snq.ctx, ent.OpQueryExist)
	switch _, err := snq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (snq *StatusNamespaceQuery) ExistX(ctx context.Context) bool {
	exist, err := snq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StatusNamespaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (snq *StatusNamespaceQuery) Clone() *StatusNamespaceQuery {
	if snq == nil {
		return nil
	}
	return &StatusNamespaceQuery{
		config:       snq.config,
		ctx:          snq.ctx.Clone(),
		order:        append([]statusnamespace.OrderOption{}, snq.order...),
		inters:       append([]Interceptor{}, snq.inters...),
		predicates:   append([]predicate.StatusNamespace{}, snq.predicates...),
		withStatuses: snq.withStatuses.Clone(),
		// clone intermediate query.
		sql:  snq.sql.Clone(),
		path: snq.path,
	}
}

// WithStatuses tells the query-builder to eager-load the nodes that are connected to
// the "statuses" edge. The optional arguments are used to configure the query builder of the edge.
func (snq *StatusNamespaceQuery) WithStatuses(opts ...func(*StatusQuery)) *StatusNamespaceQuery {
	query := (&StatusClient{config: snq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	snq.withStatuses = query
	return snq
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
//	client.StatusNamespace.Query().
//		GroupBy(statusnamespace.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (snq *StatusNamespaceQuery) GroupBy(field string, fields ...string) *StatusNamespaceGroupBy {
	snq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &StatusNamespaceGroupBy{build: snq}
	grbuild.flds = &snq.ctx.Fields
	grbuild.label = statusnamespace.Label
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
//	client.StatusNamespace.Query().
//		Select(statusnamespace.FieldCreatedAt).
//		Scan(ctx, &v)
func (snq *StatusNamespaceQuery) Select(fields ...string) *StatusNamespaceSelect {
	snq.ctx.Fields = append(snq.ctx.Fields, fields...)
	sbuild := &StatusNamespaceSelect{StatusNamespaceQuery: snq}
	sbuild.label = statusnamespace.Label
	sbuild.flds, sbuild.scan = &snq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a StatusNamespaceSelect configured with the given aggregations.
func (snq *StatusNamespaceQuery) Aggregate(fns ...AggregateFunc) *StatusNamespaceSelect {
	return snq.Select().Aggregate(fns...)
}

func (snq *StatusNamespaceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range snq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, snq); err != nil {
				return err
			}
		}
	}
	for _, f := range snq.ctx.Fields {
		if !statusnamespace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if snq.path != nil {
		prev, err := snq.path(ctx)
		if err != nil {
			return err
		}
		snq.sql = prev
	}
	return nil
}

func (snq *StatusNamespaceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*StatusNamespace, error) {
	var (
		nodes       = []*StatusNamespace{}
		_spec       = snq.querySpec()
		loadedTypes = [1]bool{
			snq.withStatuses != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*StatusNamespace).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &StatusNamespace{config: snq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(snq.modifiers) > 0 {
		_spec.Modifiers = snq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, snq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := snq.withStatuses; query != nil {
		if err := snq.loadStatuses(ctx, query, nodes,
			func(n *StatusNamespace) { n.Edges.Statuses = []*Status{} },
			func(n *StatusNamespace, e *Status) { n.Edges.Statuses = append(n.Edges.Statuses, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range snq.withNamedStatuses {
		if err := snq.loadStatuses(ctx, query, nodes,
			func(n *StatusNamespace) { n.appendNamedStatuses(name) },
			func(n *StatusNamespace, e *Status) { n.appendNamedStatuses(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range snq.loadTotal {
		if err := snq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (snq *StatusNamespaceQuery) loadStatuses(ctx context.Context, query *StatusQuery, nodes []*StatusNamespace, init func(*StatusNamespace), assign func(*StatusNamespace, *Status)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID]*StatusNamespace)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(status.FieldStatusNamespaceID)
	}
	query.Where(predicate.Status(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(statusnamespace.StatusesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.StatusNamespaceID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "status_namespace_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (snq *StatusNamespaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := snq.querySpec()
	if len(snq.modifiers) > 0 {
		_spec.Modifiers = snq.modifiers
	}
	_spec.Node.Columns = snq.ctx.Fields
	if len(snq.ctx.Fields) > 0 {
		_spec.Unique = snq.ctx.Unique != nil && *snq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, snq.driver, _spec)
}

func (snq *StatusNamespaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(statusnamespace.Table, statusnamespace.Columns, sqlgraph.NewFieldSpec(statusnamespace.FieldID, field.TypeString))
	_spec.From = snq.sql
	if unique := snq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if snq.path != nil {
		_spec.Unique = true
	}
	if fields := snq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, statusnamespace.FieldID)
		for i := range fields {
			if fields[i] != statusnamespace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := snq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := snq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := snq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := snq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (snq *StatusNamespaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(snq.driver.Dialect())
	t1 := builder.Table(statusnamespace.Table)
	columns := snq.ctx.Fields
	if len(columns) == 0 {
		columns = statusnamespace.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if snq.sql != nil {
		selector = snq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if snq.ctx.Unique != nil && *snq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range snq.predicates {
		p(selector)
	}
	for _, p := range snq.order {
		p(selector)
	}
	if offset := snq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := snq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedStatuses tells the query-builder to eager-load the nodes that are connected to the "statuses"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (snq *StatusNamespaceQuery) WithNamedStatuses(name string, opts ...func(*StatusQuery)) *StatusNamespaceQuery {
	query := (&StatusClient{config: snq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if snq.withNamedStatuses == nil {
		snq.withNamedStatuses = make(map[string]*StatusQuery)
	}
	snq.withNamedStatuses[name] = query
	return snq
}

// StatusNamespaceGroupBy is the group-by builder for StatusNamespace entities.
type StatusNamespaceGroupBy struct {
	selector
	build *StatusNamespaceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sngb *StatusNamespaceGroupBy) Aggregate(fns ...AggregateFunc) *StatusNamespaceGroupBy {
	sngb.fns = append(sngb.fns, fns...)
	return sngb
}

// Scan applies the selector query and scans the result into the given value.
func (sngb *StatusNamespaceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sngb.build.ctx, ent.OpQueryGroupBy)
	if err := sngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StatusNamespaceQuery, *StatusNamespaceGroupBy](ctx, sngb.build, sngb, sngb.build.inters, v)
}

func (sngb *StatusNamespaceGroupBy) sqlScan(ctx context.Context, root *StatusNamespaceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sngb.fns))
	for _, fn := range sngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sngb.flds)+len(sngb.fns))
		for _, f := range *sngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// StatusNamespaceSelect is the builder for selecting fields of StatusNamespace entities.
type StatusNamespaceSelect struct {
	*StatusNamespaceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sns *StatusNamespaceSelect) Aggregate(fns ...AggregateFunc) *StatusNamespaceSelect {
	sns.fns = append(sns.fns, fns...)
	return sns
}

// Scan applies the selector query and scans the result into the given value.
func (sns *StatusNamespaceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sns.ctx, ent.OpQuerySelect)
	if err := sns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StatusNamespaceQuery, *StatusNamespaceSelect](ctx, sns.StatusNamespaceQuery, sns, sns.inters, v)
}

func (sns *StatusNamespaceSelect) sqlScan(ctx context.Context, root *StatusNamespaceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sns.fns))
	for _, fn := range sns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
