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

package annotation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/x/gidx"
)

const (
	// Label holds the string label denoting the annotation type in the database.
	Label = "annotation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldMetadataID holds the string denoting the metadata_id field in the database.
	FieldMetadataID = "metadata_id"
	// FieldAnnotationNamespaceID holds the string denoting the annotation_namespace_id field in the database.
	FieldAnnotationNamespaceID = "annotation_namespace_id"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "json_data"
	// EdgeNamespace holds the string denoting the namespace edge name in mutations.
	EdgeNamespace = "namespace"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// Table holds the table name of the annotation in the database.
	Table = "annotations"
	// NamespaceTable is the table that holds the namespace relation/edge.
	NamespaceTable = "annotations"
	// NamespaceInverseTable is the table name for the AnnotationNamespace entity.
	// It exists in this package in order to avoid circular dependency with the "annotationnamespace" package.
	NamespaceInverseTable = "annotation_namespaces"
	// NamespaceColumn is the table column denoting the namespace relation/edge.
	NamespaceColumn = "annotation_namespace_id"
	// MetadataTable is the table that holds the metadata relation/edge.
	MetadataTable = "annotations"
	// MetadataInverseTable is the table name for the Metadata entity.
	// It exists in this package in order to avoid circular dependency with the "metadata" package.
	MetadataInverseTable = "metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "metadata_id"
)

// Columns holds all SQL columns for annotation fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldMetadataID,
	FieldAnnotationNamespaceID,
	FieldData,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() gidx.PrefixedID
)

// OrderOption defines the ordering options for the Annotation queries.
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

// ByMetadataID orders the results by the metadata_id field.
func ByMetadataID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMetadataID, opts...).ToFunc()
}

// ByAnnotationNamespaceID orders the results by the annotation_namespace_id field.
func ByAnnotationNamespaceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAnnotationNamespaceID, opts...).ToFunc()
}

// ByNamespaceField orders the results by namespace field.
func ByNamespaceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNamespaceStep(), sql.OrderByField(field, opts...))
	}
}

// ByMetadataField orders the results by metadata field.
func ByMetadataField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMetadataStep(), sql.OrderByField(field, opts...))
	}
}
func newNamespaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NamespaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, NamespaceTable, NamespaceColumn),
	)
}
func newMetadataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MetadataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, MetadataTable, MetadataColumn),
	)
}
