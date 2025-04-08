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
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotation"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotationnamespace"
	"go.infratographer.com/metadata-api/internal/ent/generated/metadata"
	"go.infratographer.com/metadata-api/internal/ent/generated/status"
	"go.infratographer.com/metadata-api/internal/ent/generated/statusnamespace"
	"go.infratographer.com/x/gidx"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

var annotationImplementors = []string{"Annotation", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Annotation) IsNode() {}

var annotationnamespaceImplementors = []string{"AnnotationNamespace", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*AnnotationNamespace) IsNode() {}

var metadataImplementors = []string{"Metadata", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Metadata) IsNode() {}

var statusImplementors = []string{"Status", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Status) IsNode() {}

var statusnamespaceImplementors = []string{"StatusNamespace", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*StatusNamespace) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, gidx.PrefixedID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, gidx.PrefixedID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, gidx.PrefixedID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id gidx.PrefixedID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id gidx.PrefixedID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id gidx.PrefixedID) (Noder, error) {
	switch table {
	case annotation.Table:
		var uid gidx.PrefixedID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Annotation.Query().
			Where(annotation.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, annotationImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case annotationnamespace.Table:
		var uid gidx.PrefixedID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.AnnotationNamespace.Query().
			Where(annotationnamespace.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, annotationnamespaceImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case metadata.Table:
		var uid gidx.PrefixedID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Metadata.Query().
			Where(metadata.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, metadataImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case status.Table:
		var uid gidx.PrefixedID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Status.Query().
			Where(status.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, statusImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case statusnamespace.Table:
		var uid gidx.PrefixedID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.StatusNamespace.Query().
			Where(statusnamespace.ID(uid))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, statusnamespaceImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []gidx.PrefixedID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]gidx.PrefixedID)
	id2idx := make(map[gidx.PrefixedID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []gidx.PrefixedID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[gidx.PrefixedID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case annotation.Table:
		query := c.Annotation.Query().
			Where(annotation.IDIn(ids...))
		query, err := query.CollectFields(ctx, annotationImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case annotationnamespace.Table:
		query := c.AnnotationNamespace.Query().
			Where(annotationnamespace.IDIn(ids...))
		query, err := query.CollectFields(ctx, annotationnamespaceImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case metadata.Table:
		query := c.Metadata.Query().
			Where(metadata.IDIn(ids...))
		query, err := query.CollectFields(ctx, metadataImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case status.Table:
		query := c.Status.Query().
			Where(status.IDIn(ids...))
		query, err := query.CollectFields(ctx, statusImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case statusnamespace.Table:
		query := c.StatusNamespace.Query().
			Where(statusnamespace.IDIn(ids...))
		query, err := query.CollectFields(ctx, statusnamespaceImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
