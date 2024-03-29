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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/metadata-api/internal/ent/generated/predicate"
	"go.infratographer.com/metadata-api/internal/ent/generated/statusnamespace"
)

// StatusNamespaceDelete is the builder for deleting a StatusNamespace entity.
type StatusNamespaceDelete struct {
	config
	hooks    []Hook
	mutation *StatusNamespaceMutation
}

// Where appends a list predicates to the StatusNamespaceDelete builder.
func (snd *StatusNamespaceDelete) Where(ps ...predicate.StatusNamespace) *StatusNamespaceDelete {
	snd.mutation.Where(ps...)
	return snd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (snd *StatusNamespaceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, snd.sqlExec, snd.mutation, snd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (snd *StatusNamespaceDelete) ExecX(ctx context.Context) int {
	n, err := snd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (snd *StatusNamespaceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(statusnamespace.Table, sqlgraph.NewFieldSpec(statusnamespace.FieldID, field.TypeString))
	if ps := snd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, snd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	snd.mutation.done = true
	return affected, err
}

// StatusNamespaceDeleteOne is the builder for deleting a single StatusNamespace entity.
type StatusNamespaceDeleteOne struct {
	snd *StatusNamespaceDelete
}

// Where appends a list predicates to the StatusNamespaceDelete builder.
func (sndo *StatusNamespaceDeleteOne) Where(ps ...predicate.StatusNamespace) *StatusNamespaceDeleteOne {
	sndo.snd.mutation.Where(ps...)
	return sndo
}

// Exec executes the deletion query.
func (sndo *StatusNamespaceDeleteOne) Exec(ctx context.Context) error {
	n, err := sndo.snd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{statusnamespace.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sndo *StatusNamespaceDeleteOne) ExecX(ctx context.Context) {
	if err := sndo.Exec(ctx); err != nil {
		panic(err)
	}
}
