// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotation"
	"go.infratographer.com/metadata-api/internal/ent/generated/metadata"
	"go.infratographer.com/metadata-api/internal/ent/generated/status"
	"go.infratographer.com/x/gidx"
)

// MetadataCreate is the builder for creating a Metadata entity.
type MetadataCreate struct {
	config
	mutation *MetadataMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MetadataCreate) SetCreatedAt(t time.Time) *MetadataCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MetadataCreate) SetNillableCreatedAt(t *time.Time) *MetadataCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MetadataCreate) SetUpdatedAt(t time.Time) *MetadataCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MetadataCreate) SetNillableUpdatedAt(t *time.Time) *MetadataCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetNodeID sets the "node_id" field.
func (mc *MetadataCreate) SetNodeID(gi gidx.PrefixedID) *MetadataCreate {
	mc.mutation.SetNodeID(gi)
	return mc
}

// SetID sets the "id" field.
func (mc *MetadataCreate) SetID(gi gidx.PrefixedID) *MetadataCreate {
	mc.mutation.SetID(gi)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MetadataCreate) SetNillableID(gi *gidx.PrefixedID) *MetadataCreate {
	if gi != nil {
		mc.SetID(*gi)
	}
	return mc
}

// AddAnnotationIDs adds the "annotations" edge to the Annotation entity by IDs.
func (mc *MetadataCreate) AddAnnotationIDs(ids ...gidx.PrefixedID) *MetadataCreate {
	mc.mutation.AddAnnotationIDs(ids...)
	return mc
}

// AddAnnotations adds the "annotations" edges to the Annotation entity.
func (mc *MetadataCreate) AddAnnotations(a ...*Annotation) *MetadataCreate {
	ids := make([]gidx.PrefixedID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mc.AddAnnotationIDs(ids...)
}

// AddStatusIDs adds the "statuses" edge to the Status entity by IDs.
func (mc *MetadataCreate) AddStatusIDs(ids ...gidx.PrefixedID) *MetadataCreate {
	mc.mutation.AddStatusIDs(ids...)
	return mc
}

// AddStatuses adds the "statuses" edges to the Status entity.
func (mc *MetadataCreate) AddStatuses(s ...*Status) *MetadataCreate {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mc.AddStatusIDs(ids...)
}

// Mutation returns the MetadataMutation object of the builder.
func (mc *MetadataCreate) Mutation() *MetadataMutation {
	return mc.mutation
}

// Save creates the Metadata in the database.
func (mc *MetadataCreate) Save(ctx context.Context) (*Metadata, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MetadataCreate) SaveX(ctx context.Context) *Metadata {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MetadataCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MetadataCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MetadataCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := metadata.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := metadata.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := metadata.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MetadataCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "Metadata.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "Metadata.updated_at"`)}
	}
	if _, ok := mc.mutation.NodeID(); !ok {
		return &ValidationError{Name: "node_id", err: errors.New(`generated: missing required field "Metadata.node_id"`)}
	}
	return nil
}

func (mc *MetadataCreate) sqlSave(ctx context.Context) (*Metadata, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*gidx.PrefixedID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MetadataCreate) createSpec() (*Metadata, *sqlgraph.CreateSpec) {
	var (
		_node = &Metadata{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(metadata.Table, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(metadata.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(metadata.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.NodeID(); ok {
		_spec.SetField(metadata.FieldNodeID, field.TypeString, value)
		_node.NodeID = value
	}
	if nodes := mc.mutation.AnnotationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.AnnotationsTable,
			Columns: []string{metadata.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.StatusesTable,
			Columns: []string{metadata.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MetadataCreateBulk is the builder for creating many Metadata entities in bulk.
type MetadataCreateBulk struct {
	config
	builders []*MetadataCreate
}

// Save creates the Metadata entities in the database.
func (mcb *MetadataCreateBulk) Save(ctx context.Context) ([]*Metadata, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Metadata, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetadataMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MetadataCreateBulk) SaveX(ctx context.Context) []*Metadata {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MetadataCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MetadataCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
