// Code generated by entc, DO NOT EDIT.

package eventhooks

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"go.infratographer.com/metadata-api/internal/ent/generated"
	"go.infratographer.com/metadata-api/internal/ent/generated/hook"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/gidx"
)

func AnnotationHooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.AnnotationFunc(func(ctx context.Context, m *generated.AnnotationMutation) (ent.Value, error) {
					var err error
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					changeset := []events.FieldChange{}
					cv_created_at := ""
					created_at, ok := m.CreatedAt()

					if ok {
						cv_created_at = created_at.Format(time.RFC3339)
						pv_created_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldCreatedAt(ctx)
							if err != nil {
								pv_created_at = "<unknown>"
							} else {
								pv_created_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "created_at",
							PreviousValue: pv_created_at,
							CurrentValue:  cv_created_at,
						})
					}

					cv_updated_at := ""
					updated_at, ok := m.UpdatedAt()

					if ok {
						cv_updated_at = updated_at.Format(time.RFC3339)
						pv_updated_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldUpdatedAt(ctx)
							if err != nil {
								pv_updated_at = "<unknown>"
							} else {
								pv_updated_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "updated_at",
							PreviousValue: pv_updated_at,
							CurrentValue:  cv_updated_at,
						})
					}

					cv_metadata_id := ""
					metadata_id, ok := m.MetadataID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						metadata_id, err = m.OldMetadataID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, metadata_id)

					if ok {
						cv_metadata_id = fmt.Sprintf("%s", fmt.Sprint(metadata_id))
						pv_metadata_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldMetadataID(ctx)
							if err != nil {
								pv_metadata_id = "<unknown>"
							} else {
								pv_metadata_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "metadata_id",
							PreviousValue: pv_metadata_id,
							CurrentValue:  cv_metadata_id,
						})
					}

					cv_annotation_namespace_id := ""
					annotation_namespace_id, ok := m.AnnotationNamespaceID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						annotation_namespace_id, err = m.OldAnnotationNamespaceID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, annotation_namespace_id)

					if ok {
						cv_annotation_namespace_id = fmt.Sprintf("%s", fmt.Sprint(annotation_namespace_id))
						pv_annotation_namespace_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldAnnotationNamespaceID(ctx)
							if err != nil {
								pv_annotation_namespace_id = "<unknown>"
							} else {
								pv_annotation_namespace_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "annotation_namespace_id",
							PreviousValue: pv_annotation_namespace_id,
							CurrentValue:  cv_annotation_namespace_id,
						})
					}

					cv_data := ""
					data, ok := m.Data()

					if ok {
						cv_data = fmt.Sprintf("%s", fmt.Sprint(data))
						pv_data := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldData(ctx)
							if err != nil {
								pv_data = "<unknown>"
							} else {
								pv_data = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "data",
							PreviousValue: pv_data,
							CurrentValue:  cv_data,
						})
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
						FieldChanges:         changeset,
					}

					// complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					if err := m.EventsPublisher.PublishChange(ctx, "annotation", msg); err != nil {
						return nil, fmt.Errorf("failed to publish change: %w", err)
					}

					return retValue, nil
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),

		// Delete Hook
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.AnnotationFunc(func(ctx context.Context, m *generated.AnnotationMutation) (ent.Value, error) {
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					dbObj, err := m.Client().Annotation.Get(ctx, objID)
					if err != nil {
						return nil, fmt.Errorf("failed to load object to get values for event, err %w", err)
					}

					additionalSubjects = append(additionalSubjects, dbObj.MetadataID)

					additionalSubjects = append(additionalSubjects, dbObj.AnnotationNamespaceID)

					// we have all the info we need, now complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
					}

					if err := m.EventsPublisher.PublishChange(ctx, "annotation", msg); err != nil {
						return nil, fmt.Errorf("failed to publish change: %w", err)
					}

					return retValue, nil
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}
func AnnotationNamespaceHooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.AnnotationNamespaceFunc(func(ctx context.Context, m *generated.AnnotationNamespaceMutation) (ent.Value, error) {
					var err error
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					changeset := []events.FieldChange{}
					cv_created_at := ""
					created_at, ok := m.CreatedAt()

					if ok {
						cv_created_at = created_at.Format(time.RFC3339)
						pv_created_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldCreatedAt(ctx)
							if err != nil {
								pv_created_at = "<unknown>"
							} else {
								pv_created_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "created_at",
							PreviousValue: pv_created_at,
							CurrentValue:  cv_created_at,
						})
					}

					cv_updated_at := ""
					updated_at, ok := m.UpdatedAt()

					if ok {
						cv_updated_at = updated_at.Format(time.RFC3339)
						pv_updated_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldUpdatedAt(ctx)
							if err != nil {
								pv_updated_at = "<unknown>"
							} else {
								pv_updated_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "updated_at",
							PreviousValue: pv_updated_at,
							CurrentValue:  cv_updated_at,
						})
					}

					cv_name := ""
					name, ok := m.Name()

					if ok {
						cv_name = fmt.Sprintf("%s", fmt.Sprint(name))
						pv_name := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldName(ctx)
							if err != nil {
								pv_name = "<unknown>"
							} else {
								pv_name = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "name",
							PreviousValue: pv_name,
							CurrentValue:  cv_name,
						})
					}

					cv_owner_id := ""
					owner_id, ok := m.OwnerID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						owner_id, err = m.OldOwnerID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, owner_id)

					if ok {
						cv_owner_id = fmt.Sprintf("%s", fmt.Sprint(owner_id))
						pv_owner_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldOwnerID(ctx)
							if err != nil {
								pv_owner_id = "<unknown>"
							} else {
								pv_owner_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "owner_id",
							PreviousValue: pv_owner_id,
							CurrentValue:  cv_owner_id,
						})
					}

					cv_private := ""
					private, ok := m.Private()

					if ok {
						cv_private = fmt.Sprintf("%s", fmt.Sprint(private))
						pv_private := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldPrivate(ctx)
							if err != nil {
								pv_private = "<unknown>"
							} else {
								pv_private = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "private",
							PreviousValue: pv_private,
							CurrentValue:  cv_private,
						})
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
						FieldChanges:         changeset,
					}

					// complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					if err := m.EventsPublisher.PublishChange(ctx, "annotation-namespace", msg); err != nil {
						return nil, fmt.Errorf("failed to publish change: %w", err)
					}

					return retValue, nil
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),

		// Delete Hook
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.AnnotationNamespaceFunc(func(ctx context.Context, m *generated.AnnotationNamespaceMutation) (ent.Value, error) {
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					dbObj, err := m.Client().AnnotationNamespace.Get(ctx, objID)
					if err != nil {
						return nil, fmt.Errorf("failed to load object to get values for event, err %w", err)
					}

					additionalSubjects = append(additionalSubjects, dbObj.OwnerID)

					// we have all the info we need, now complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
					}

					if err := m.EventsPublisher.PublishChange(ctx, "annotation-namespace", msg); err != nil {
						return nil, fmt.Errorf("failed to publish change: %w", err)
					}

					return retValue, nil
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}
func MetadataHooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.MetadataFunc(func(ctx context.Context, m *generated.MetadataMutation) (ent.Value, error) {
					var err error
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					changeset := []events.FieldChange{}
					cv_created_at := ""
					created_at, ok := m.CreatedAt()

					if ok {
						cv_created_at = created_at.Format(time.RFC3339)
						pv_created_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldCreatedAt(ctx)
							if err != nil {
								pv_created_at = "<unknown>"
							} else {
								pv_created_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "created_at",
							PreviousValue: pv_created_at,
							CurrentValue:  cv_created_at,
						})
					}

					cv_updated_at := ""
					updated_at, ok := m.UpdatedAt()

					if ok {
						cv_updated_at = updated_at.Format(time.RFC3339)
						pv_updated_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldUpdatedAt(ctx)
							if err != nil {
								pv_updated_at = "<unknown>"
							} else {
								pv_updated_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "updated_at",
							PreviousValue: pv_updated_at,
							CurrentValue:  cv_updated_at,
						})
					}

					cv_node_id := ""
					node_id, ok := m.NodeID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						node_id, err = m.OldNodeID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, node_id)

					if ok {
						cv_node_id = fmt.Sprintf("%s", fmt.Sprint(node_id))
						pv_node_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldNodeID(ctx)
							if err != nil {
								pv_node_id = "<unknown>"
							} else {
								pv_node_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "node_id",
							PreviousValue: pv_node_id,
							CurrentValue:  cv_node_id,
						})
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
						FieldChanges:         changeset,
					}

					// complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					if err := m.EventsPublisher.PublishChange(ctx, "metadata", msg); err != nil {
						return nil, fmt.Errorf("failed to publish change: %w", err)
					}

					return retValue, nil
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),

		// Delete Hook
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.MetadataFunc(func(ctx context.Context, m *generated.MetadataMutation) (ent.Value, error) {
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					dbObj, err := m.Client().Metadata.Get(ctx, objID)
					if err != nil {
						return nil, fmt.Errorf("failed to load object to get values for event, err %w", err)
					}

					additionalSubjects = append(additionalSubjects, dbObj.NodeID)

					// we have all the info we need, now complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
					}

					if err := m.EventsPublisher.PublishChange(ctx, "metadata", msg); err != nil {
						return nil, fmt.Errorf("failed to publish change: %w", err)
					}

					return retValue, nil
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}
func StatusHooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.StatusFunc(func(ctx context.Context, m *generated.StatusMutation) (ent.Value, error) {
					var err error
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					changeset := []events.FieldChange{}
					cv_created_at := ""
					created_at, ok := m.CreatedAt()

					if ok {
						cv_created_at = created_at.Format(time.RFC3339)
						pv_created_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldCreatedAt(ctx)
							if err != nil {
								pv_created_at = "<unknown>"
							} else {
								pv_created_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "created_at",
							PreviousValue: pv_created_at,
							CurrentValue:  cv_created_at,
						})
					}

					cv_updated_at := ""
					updated_at, ok := m.UpdatedAt()

					if ok {
						cv_updated_at = updated_at.Format(time.RFC3339)
						pv_updated_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldUpdatedAt(ctx)
							if err != nil {
								pv_updated_at = "<unknown>"
							} else {
								pv_updated_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "updated_at",
							PreviousValue: pv_updated_at,
							CurrentValue:  cv_updated_at,
						})
					}

					cv_metadata_id := ""
					metadata_id, ok := m.MetadataID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						metadata_id, err = m.OldMetadataID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, metadata_id)

					if ok {
						cv_metadata_id = fmt.Sprintf("%s", fmt.Sprint(metadata_id))
						pv_metadata_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldMetadataID(ctx)
							if err != nil {
								pv_metadata_id = "<unknown>"
							} else {
								pv_metadata_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "metadata_id",
							PreviousValue: pv_metadata_id,
							CurrentValue:  cv_metadata_id,
						})
					}

					cv_status_namespace_id := ""
					status_namespace_id, ok := m.StatusNamespaceID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						status_namespace_id, err = m.OldStatusNamespaceID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, status_namespace_id)

					if ok {
						cv_status_namespace_id = fmt.Sprintf("%s", fmt.Sprint(status_namespace_id))
						pv_status_namespace_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldStatusNamespaceID(ctx)
							if err != nil {
								pv_status_namespace_id = "<unknown>"
							} else {
								pv_status_namespace_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "status_namespace_id",
							PreviousValue: pv_status_namespace_id,
							CurrentValue:  cv_status_namespace_id,
						})
					}

					cv_source := ""
					source, ok := m.Source()

					if ok {
						cv_source = fmt.Sprintf("%s", fmt.Sprint(source))
						pv_source := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldSource(ctx)
							if err != nil {
								pv_source = "<unknown>"
							} else {
								pv_source = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "source",
							PreviousValue: pv_source,
							CurrentValue:  cv_source,
						})
					}

					cv_data := ""
					data, ok := m.Data()

					if ok {
						cv_data = fmt.Sprintf("%s", fmt.Sprint(data))
						pv_data := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldData(ctx)
							if err != nil {
								pv_data = "<unknown>"
							} else {
								pv_data = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "data",
							PreviousValue: pv_data,
							CurrentValue:  cv_data,
						})
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
						FieldChanges:         changeset,
					}

					// complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					if err := m.EventsPublisher.PublishChange(ctx, "status", msg); err != nil {
						return nil, fmt.Errorf("failed to publish change: %w", err)
					}

					return retValue, nil
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),

		// Delete Hook
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.StatusFunc(func(ctx context.Context, m *generated.StatusMutation) (ent.Value, error) {
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					dbObj, err := m.Client().Status.Get(ctx, objID)
					if err != nil {
						return nil, fmt.Errorf("failed to load object to get values for event, err %w", err)
					}

					additionalSubjects = append(additionalSubjects, dbObj.MetadataID)

					additionalSubjects = append(additionalSubjects, dbObj.StatusNamespaceID)

					// we have all the info we need, now complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
					}

					if err := m.EventsPublisher.PublishChange(ctx, "status", msg); err != nil {
						return nil, fmt.Errorf("failed to publish change: %w", err)
					}

					return retValue, nil
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}
func StatusNamespaceHooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.StatusNamespaceFunc(func(ctx context.Context, m *generated.StatusNamespaceMutation) (ent.Value, error) {
					var err error
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					changeset := []events.FieldChange{}
					cv_created_at := ""
					created_at, ok := m.CreatedAt()

					if ok {
						cv_created_at = created_at.Format(time.RFC3339)
						pv_created_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldCreatedAt(ctx)
							if err != nil {
								pv_created_at = "<unknown>"
							} else {
								pv_created_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "created_at",
							PreviousValue: pv_created_at,
							CurrentValue:  cv_created_at,
						})
					}

					cv_updated_at := ""
					updated_at, ok := m.UpdatedAt()

					if ok {
						cv_updated_at = updated_at.Format(time.RFC3339)
						pv_updated_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldUpdatedAt(ctx)
							if err != nil {
								pv_updated_at = "<unknown>"
							} else {
								pv_updated_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "updated_at",
							PreviousValue: pv_updated_at,
							CurrentValue:  cv_updated_at,
						})
					}

					cv_name := ""
					name, ok := m.Name()

					if ok {
						cv_name = fmt.Sprintf("%s", fmt.Sprint(name))
						pv_name := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldName(ctx)
							if err != nil {
								pv_name = "<unknown>"
							} else {
								pv_name = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "name",
							PreviousValue: pv_name,
							CurrentValue:  cv_name,
						})
					}

					cv_resource_provider_id := ""
					resource_provider_id, ok := m.ResourceProviderID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						resource_provider_id, err = m.OldResourceProviderID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, resource_provider_id)

					if ok {
						cv_resource_provider_id = fmt.Sprintf("%s", fmt.Sprint(resource_provider_id))
						pv_resource_provider_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldResourceProviderID(ctx)
							if err != nil {
								pv_resource_provider_id = "<unknown>"
							} else {
								pv_resource_provider_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "resource_provider_id",
							PreviousValue: pv_resource_provider_id,
							CurrentValue:  cv_resource_provider_id,
						})
					}

					cv_private := ""
					private, ok := m.Private()

					if ok {
						cv_private = fmt.Sprintf("%s", fmt.Sprint(private))
						pv_private := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldPrivate(ctx)
							if err != nil {
								pv_private = "<unknown>"
							} else {
								pv_private = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "private",
							PreviousValue: pv_private,
							CurrentValue:  cv_private,
						})
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
						FieldChanges:         changeset,
					}

					// complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					if err := m.EventsPublisher.PublishChange(ctx, "status-namespace", msg); err != nil {
						return nil, fmt.Errorf("failed to publish change: %w", err)
					}

					return retValue, nil
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),

		// Delete Hook
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.StatusNamespaceFunc(func(ctx context.Context, m *generated.StatusNamespaceMutation) (ent.Value, error) {
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					dbObj, err := m.Client().StatusNamespace.Get(ctx, objID)
					if err != nil {
						return nil, fmt.Errorf("failed to load object to get values for event, err %w", err)
					}

					additionalSubjects = append(additionalSubjects, dbObj.ResourceProviderID)

					// we have all the info we need, now complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
					}

					if err := m.EventsPublisher.PublishChange(ctx, "status-namespace", msg); err != nil {
						return nil, fmt.Errorf("failed to publish change: %w", err)
					}

					return retValue, nil
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}

func EventHooks(c *generated.Client) {
	c.Annotation.Use(AnnotationHooks()...)

	c.AnnotationNamespace.Use(AnnotationNamespaceHooks()...)

	c.Metadata.Use(MetadataHooks()...)

	c.Status.Use(StatusHooks()...)

	c.StatusNamespace.Use(StatusNamespaceHooks()...)

}

func eventType(op ent.Op) string {
	switch op {
	case ent.OpCreate:
		return string(events.CreateChangeType)
	case ent.OpUpdate, ent.OpUpdateOne:
		return string(events.UpdateChangeType)
	case ent.OpDelete, ent.OpDeleteOne:
		return string(events.DeleteChangeType)
	default:
		return "unknown"
	}
}
