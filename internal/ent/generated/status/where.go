// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package status

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/metadata-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// ID filters vertices based on their ID field.
func ID(id gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldUpdatedAt, v))
}

// MetadataID applies equality check predicate on the "metadata_id" field. It's identical to MetadataIDEQ.
func MetadataID(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldMetadataID, v))
}

// StatusNamespaceID applies equality check predicate on the "status_namespace_id" field. It's identical to StatusNamespaceIDEQ.
func StatusNamespaceID(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldStatusNamespaceID, v))
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldSource, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldLTE(FieldUpdatedAt, v))
}

// MetadataIDEQ applies the EQ predicate on the "metadata_id" field.
func MetadataIDEQ(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldMetadataID, v))
}

// MetadataIDNEQ applies the NEQ predicate on the "metadata_id" field.
func MetadataIDNEQ(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldMetadataID, v))
}

// MetadataIDIn applies the In predicate on the "metadata_id" field.
func MetadataIDIn(vs ...gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldMetadataID, vs...))
}

// MetadataIDNotIn applies the NotIn predicate on the "metadata_id" field.
func MetadataIDNotIn(vs ...gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldMetadataID, vs...))
}

// MetadataIDGT applies the GT predicate on the "metadata_id" field.
func MetadataIDGT(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldGT(FieldMetadataID, v))
}

// MetadataIDGTE applies the GTE predicate on the "metadata_id" field.
func MetadataIDGTE(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldGTE(FieldMetadataID, v))
}

// MetadataIDLT applies the LT predicate on the "metadata_id" field.
func MetadataIDLT(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldLT(FieldMetadataID, v))
}

// MetadataIDLTE applies the LTE predicate on the "metadata_id" field.
func MetadataIDLTE(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldLTE(FieldMetadataID, v))
}

// MetadataIDContains applies the Contains predicate on the "metadata_id" field.
func MetadataIDContains(v gidx.PrefixedID) predicate.Status {
	vc := string(v)
	return predicate.Status(sql.FieldContains(FieldMetadataID, vc))
}

// MetadataIDHasPrefix applies the HasPrefix predicate on the "metadata_id" field.
func MetadataIDHasPrefix(v gidx.PrefixedID) predicate.Status {
	vc := string(v)
	return predicate.Status(sql.FieldHasPrefix(FieldMetadataID, vc))
}

// MetadataIDHasSuffix applies the HasSuffix predicate on the "metadata_id" field.
func MetadataIDHasSuffix(v gidx.PrefixedID) predicate.Status {
	vc := string(v)
	return predicate.Status(sql.FieldHasSuffix(FieldMetadataID, vc))
}

// MetadataIDEqualFold applies the EqualFold predicate on the "metadata_id" field.
func MetadataIDEqualFold(v gidx.PrefixedID) predicate.Status {
	vc := string(v)
	return predicate.Status(sql.FieldEqualFold(FieldMetadataID, vc))
}

// MetadataIDContainsFold applies the ContainsFold predicate on the "metadata_id" field.
func MetadataIDContainsFold(v gidx.PrefixedID) predicate.Status {
	vc := string(v)
	return predicate.Status(sql.FieldContainsFold(FieldMetadataID, vc))
}

// StatusNamespaceIDEQ applies the EQ predicate on the "status_namespace_id" field.
func StatusNamespaceIDEQ(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldStatusNamespaceID, v))
}

// StatusNamespaceIDNEQ applies the NEQ predicate on the "status_namespace_id" field.
func StatusNamespaceIDNEQ(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldStatusNamespaceID, v))
}

// StatusNamespaceIDIn applies the In predicate on the "status_namespace_id" field.
func StatusNamespaceIDIn(vs ...gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldStatusNamespaceID, vs...))
}

// StatusNamespaceIDNotIn applies the NotIn predicate on the "status_namespace_id" field.
func StatusNamespaceIDNotIn(vs ...gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldStatusNamespaceID, vs...))
}

// StatusNamespaceIDGT applies the GT predicate on the "status_namespace_id" field.
func StatusNamespaceIDGT(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldGT(FieldStatusNamespaceID, v))
}

// StatusNamespaceIDGTE applies the GTE predicate on the "status_namespace_id" field.
func StatusNamespaceIDGTE(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldGTE(FieldStatusNamespaceID, v))
}

// StatusNamespaceIDLT applies the LT predicate on the "status_namespace_id" field.
func StatusNamespaceIDLT(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldLT(FieldStatusNamespaceID, v))
}

// StatusNamespaceIDLTE applies the LTE predicate on the "status_namespace_id" field.
func StatusNamespaceIDLTE(v gidx.PrefixedID) predicate.Status {
	return predicate.Status(sql.FieldLTE(FieldStatusNamespaceID, v))
}

// StatusNamespaceIDContains applies the Contains predicate on the "status_namespace_id" field.
func StatusNamespaceIDContains(v gidx.PrefixedID) predicate.Status {
	vc := string(v)
	return predicate.Status(sql.FieldContains(FieldStatusNamespaceID, vc))
}

// StatusNamespaceIDHasPrefix applies the HasPrefix predicate on the "status_namespace_id" field.
func StatusNamespaceIDHasPrefix(v gidx.PrefixedID) predicate.Status {
	vc := string(v)
	return predicate.Status(sql.FieldHasPrefix(FieldStatusNamespaceID, vc))
}

// StatusNamespaceIDHasSuffix applies the HasSuffix predicate on the "status_namespace_id" field.
func StatusNamespaceIDHasSuffix(v gidx.PrefixedID) predicate.Status {
	vc := string(v)
	return predicate.Status(sql.FieldHasSuffix(FieldStatusNamespaceID, vc))
}

// StatusNamespaceIDEqualFold applies the EqualFold predicate on the "status_namespace_id" field.
func StatusNamespaceIDEqualFold(v gidx.PrefixedID) predicate.Status {
	vc := string(v)
	return predicate.Status(sql.FieldEqualFold(FieldStatusNamespaceID, vc))
}

// StatusNamespaceIDContainsFold applies the ContainsFold predicate on the "status_namespace_id" field.
func StatusNamespaceIDContainsFold(v gidx.PrefixedID) predicate.Status {
	vc := string(v)
	return predicate.Status(sql.FieldContainsFold(FieldStatusNamespaceID, vc))
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldSource, v))
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldSource, v))
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldSource, vs...))
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldSource, vs...))
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.Status {
	return predicate.Status(sql.FieldGT(FieldSource, v))
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.Status {
	return predicate.Status(sql.FieldGTE(FieldSource, v))
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.Status {
	return predicate.Status(sql.FieldLT(FieldSource, v))
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.Status {
	return predicate.Status(sql.FieldLTE(FieldSource, v))
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.Status {
	return predicate.Status(sql.FieldContains(FieldSource, v))
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.Status {
	return predicate.Status(sql.FieldHasPrefix(FieldSource, v))
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.Status {
	return predicate.Status(sql.FieldHasSuffix(FieldSource, v))
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.Status {
	return predicate.Status(sql.FieldEqualFold(FieldSource, v))
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.Status {
	return predicate.Status(sql.FieldContainsFold(FieldSource, v))
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.StatusNamespace) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := newNamespaceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.Metadata) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := newMetadataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		p(s.Not())
	})
}
