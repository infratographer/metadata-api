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
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/metadata-api/internal/ent/generated/metadata"
	"go.infratographer.com/metadata-api/internal/ent/generated/status"
	"go.infratographer.com/metadata-api/internal/ent/generated/statusnamespace"
	"go.infratographer.com/x/gidx"
)

// Status is the model entity for the Status schema.
type Status struct {
	config `json:"-"`
	// ID of the ent.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ID of the metadata of this status
	MetadataID gidx.PrefixedID `json:"metadata_id,omitempty"`
	// StatusNamespaceID holds the value of the "status_namespace_id" field.
	StatusNamespaceID gidx.PrefixedID `json:"status_namespace_id,omitempty"`
	// Source holds the value of the "source" field.
	Source string `json:"source,omitempty"`
	// JSON formatted data of this annotation.
	Data json.RawMessage `json:"data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StatusQuery when eager-loading is set.
	Edges        StatusEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StatusEdges holds the relations/edges for other nodes in the graph.
type StatusEdges struct {
	// Namespace holds the value of the namespace edge.
	Namespace *StatusNamespace `json:"namespace,omitempty"`
	// Metadata holds the value of the metadata edge.
	Metadata *Metadata `json:"metadata,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// NamespaceOrErr returns the Namespace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StatusEdges) NamespaceOrErr() (*StatusNamespace, error) {
	if e.Namespace != nil {
		return e.Namespace, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: statusnamespace.Label}
	}
	return nil, &NotLoadedError{edge: "namespace"}
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StatusEdges) MetadataOrErr() (*Metadata, error) {
	if e.Metadata != nil {
		return e.Metadata, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: metadata.Label}
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Status) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case status.FieldData:
			values[i] = new([]byte)
		case status.FieldID, status.FieldMetadataID, status.FieldStatusNamespaceID:
			values[i] = new(gidx.PrefixedID)
		case status.FieldSource:
			values[i] = new(sql.NullString)
		case status.FieldCreatedAt, status.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Status fields.
func (s *Status) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case status.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case status.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case status.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case status.FieldMetadataID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field metadata_id", values[i])
			} else if value != nil {
				s.MetadataID = *value
			}
		case status.FieldStatusNamespaceID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field status_namespace_id", values[i])
			} else if value != nil {
				s.StatusNamespaceID = *value
			}
		case status.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				s.Source = value.String
			}
		case status.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Data); err != nil {
					return fmt.Errorf("unmarshal field data: %w", err)
				}
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Status.
// This includes values selected through modifiers, order, etc.
func (s *Status) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryNamespace queries the "namespace" edge of the Status entity.
func (s *Status) QueryNamespace() *StatusNamespaceQuery {
	return NewStatusClient(s.config).QueryNamespace(s)
}

// QueryMetadata queries the "metadata" edge of the Status entity.
func (s *Status) QueryMetadata() *MetadataQuery {
	return NewStatusClient(s.config).QueryMetadata(s)
}

// Update returns a builder for updating this Status.
// Note that you need to call Status.Unwrap() before calling this method if this Status
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Status) Update() *StatusUpdateOne {
	return NewStatusClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Status entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Status) Unwrap() *Status {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("generated: Status is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Status) String() string {
	var builder strings.Builder
	builder.WriteString("Status(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("metadata_id=")
	builder.WriteString(fmt.Sprintf("%v", s.MetadataID))
	builder.WriteString(", ")
	builder.WriteString("status_namespace_id=")
	builder.WriteString(fmt.Sprintf("%v", s.StatusNamespaceID))
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(s.Source)
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", s.Data))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (s Status) IsEntity() {}

// StatusSlice is a parsable slice of Status.
type StatusSlice []*Status
