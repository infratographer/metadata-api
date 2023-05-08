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

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/metadata-api/internal/ent/generated/metadata"
	"go.infratographer.com/x/gidx"
)

// Metadata is the model entity for the Metadata schema.
type Metadata struct {
	config `json:"-"`
	// ID of the ent.
	// ID for the metadata.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ID of the node for this metadata
	NodeID gidx.PrefixedID `json:"node_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MetadataQuery when eager-loading is set.
	Edges        MetadataEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MetadataEdges holds the relations/edges for other nodes in the graph.
type MetadataEdges struct {
	// Annotations holds the value of the annotations edge.
	Annotations []*Annotation `json:"annotations,omitempty"`
	// Statuses holds the value of the statuses edge.
	Statuses []*Status `json:"statuses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedAnnotations map[string][]*Annotation
	namedStatuses    map[string][]*Status
}

// AnnotationsOrErr returns the Annotations value or an error if the edge
// was not loaded in eager-loading.
func (e MetadataEdges) AnnotationsOrErr() ([]*Annotation, error) {
	if e.loadedTypes[0] {
		return e.Annotations, nil
	}
	return nil, &NotLoadedError{edge: "annotations"}
}

// StatusesOrErr returns the Statuses value or an error if the edge
// was not loaded in eager-loading.
func (e MetadataEdges) StatusesOrErr() ([]*Status, error) {
	if e.loadedTypes[1] {
		return e.Statuses, nil
	}
	return nil, &NotLoadedError{edge: "statuses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Metadata) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case metadata.FieldID, metadata.FieldNodeID:
			values[i] = new(gidx.PrefixedID)
		case metadata.FieldCreatedAt, metadata.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Metadata fields.
func (m *Metadata) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case metadata.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case metadata.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case metadata.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case metadata.FieldNodeID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field node_id", values[i])
			} else if value != nil {
				m.NodeID = *value
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Metadata.
// This includes values selected through modifiers, order, etc.
func (m *Metadata) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryAnnotations queries the "annotations" edge of the Metadata entity.
func (m *Metadata) QueryAnnotations() *AnnotationQuery {
	return NewMetadataClient(m.config).QueryAnnotations(m)
}

// QueryStatuses queries the "statuses" edge of the Metadata entity.
func (m *Metadata) QueryStatuses() *StatusQuery {
	return NewMetadataClient(m.config).QueryStatuses(m)
}

// Update returns a builder for updating this Metadata.
// Note that you need to call Metadata.Unwrap() before calling this method if this Metadata
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Metadata) Update() *MetadataUpdateOne {
	return NewMetadataClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Metadata entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Metadata) Unwrap() *Metadata {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("generated: Metadata is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Metadata) String() string {
	var builder strings.Builder
	builder.WriteString("Metadata(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("node_id=")
	builder.WriteString(fmt.Sprintf("%v", m.NodeID))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (m Metadata) IsEntity() {}

// NamedAnnotations returns the Annotations named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Metadata) NamedAnnotations(name string) ([]*Annotation, error) {
	if m.Edges.namedAnnotations == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedAnnotations[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Metadata) appendNamedAnnotations(name string, edges ...*Annotation) {
	if m.Edges.namedAnnotations == nil {
		m.Edges.namedAnnotations = make(map[string][]*Annotation)
	}
	if len(edges) == 0 {
		m.Edges.namedAnnotations[name] = []*Annotation{}
	} else {
		m.Edges.namedAnnotations[name] = append(m.Edges.namedAnnotations[name], edges...)
	}
}

// NamedStatuses returns the Statuses named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Metadata) NamedStatuses(name string) ([]*Status, error) {
	if m.Edges.namedStatuses == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedStatuses[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Metadata) appendNamedStatuses(name string, edges ...*Status) {
	if m.Edges.namedStatuses == nil {
		m.Edges.namedStatuses = make(map[string][]*Status)
	}
	if len(edges) == 0 {
		m.Edges.namedStatuses[name] = []*Status{}
	} else {
		m.Edges.namedStatuses[name] = append(m.Edges.namedStatuses[name], edges...)
	}
}

// MetadataSlice is a parsable slice of Metadata.
type MetadataSlice []*Metadata
