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

package annotationnamespace

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/x/gidx"
)

const (
	// Label holds the string label denoting the annotationnamespace type in the database.
	Label = "annotation_namespace"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldPrivate holds the string denoting the private field in the database.
	FieldPrivate = "private"
	// EdgeAnnotations holds the string denoting the annotations edge name in mutations.
	EdgeAnnotations = "annotations"
	// Table holds the table name of the annotationnamespace in the database.
	Table = "annotation_namespaces"
	// AnnotationsTable is the table that holds the annotations relation/edge.
	AnnotationsTable = "annotations"
	// AnnotationsInverseTable is the table name for the Annotation entity.
	// It exists in this package in order to avoid circular dependency with the "annotation" package.
	AnnotationsInverseTable = "annotations"
	// AnnotationsColumn is the table column denoting the annotations relation/edge.
	AnnotationsColumn = "annotation_namespace_id"
)

// Columns holds all SQL columns for annotationnamespace fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldOwnerID,
	FieldPrivate,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// OwnerIDValidator is a validator for the "owner_id" field. It is called by the builders before save.
	OwnerIDValidator func(string) error
	// DefaultPrivate holds the default value on creation for the "private" field.
	DefaultPrivate bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() gidx.PrefixedID
)

// OrderOption defines the ordering options for the AnnotationNamespace queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByOwnerID orders the results by the owner_id field.
func ByOwnerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerID, opts...).ToFunc()
}

// ByPrivate orders the results by the private field.
func ByPrivate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrivate, opts...).ToFunc()
}

// ByAnnotationsCount orders the results by annotations count.
func ByAnnotationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAnnotationsStep(), opts...)
	}
}

// ByAnnotations orders the results by annotations terms.
func ByAnnotations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAnnotationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAnnotationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AnnotationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, AnnotationsTable, AnnotationsColumn),
	)
}
