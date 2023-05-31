// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"time"

	"go.infratographer.com/metadata-api/internal/ent/generated/annotation"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotationnamespace"
	"go.infratographer.com/metadata-api/internal/ent/generated/metadata"
	"go.infratographer.com/metadata-api/internal/ent/generated/status"
	"go.infratographer.com/metadata-api/internal/ent/generated/statusnamespace"
	"go.infratographer.com/metadata-api/internal/ent/schema"
	"go.infratographer.com/x/gidx"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	annotationMixin := schema.Annotation{}.Mixin()
	annotationMixinFields0 := annotationMixin[0].Fields()
	_ = annotationMixinFields0
	annotationFields := schema.Annotation{}.Fields()
	_ = annotationFields
	// annotationDescCreatedAt is the schema descriptor for created_at field.
	annotationDescCreatedAt := annotationMixinFields0[0].Descriptor()
	// annotation.DefaultCreatedAt holds the default value on creation for the created_at field.
	annotation.DefaultCreatedAt = annotationDescCreatedAt.Default.(func() time.Time)
	// annotationDescUpdatedAt is the schema descriptor for updated_at field.
	annotationDescUpdatedAt := annotationMixinFields0[1].Descriptor()
	// annotation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	annotation.DefaultUpdatedAt = annotationDescUpdatedAt.Default.(func() time.Time)
	// annotation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	annotation.UpdateDefaultUpdatedAt = annotationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// annotationDescID is the schema descriptor for id field.
	annotationDescID := annotationFields[0].Descriptor()
	// annotation.DefaultID holds the default value on creation for the id field.
	annotation.DefaultID = annotationDescID.Default.(func() gidx.PrefixedID)
	annotationnamespaceMixin := schema.AnnotationNamespace{}.Mixin()
	annotationnamespaceMixinFields0 := annotationnamespaceMixin[0].Fields()
	_ = annotationnamespaceMixinFields0
	annotationnamespaceFields := schema.AnnotationNamespace{}.Fields()
	_ = annotationnamespaceFields
	// annotationnamespaceDescCreatedAt is the schema descriptor for created_at field.
	annotationnamespaceDescCreatedAt := annotationnamespaceMixinFields0[0].Descriptor()
	// annotationnamespace.DefaultCreatedAt holds the default value on creation for the created_at field.
	annotationnamespace.DefaultCreatedAt = annotationnamespaceDescCreatedAt.Default.(func() time.Time)
	// annotationnamespaceDescUpdatedAt is the schema descriptor for updated_at field.
	annotationnamespaceDescUpdatedAt := annotationnamespaceMixinFields0[1].Descriptor()
	// annotationnamespace.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	annotationnamespace.DefaultUpdatedAt = annotationnamespaceDescUpdatedAt.Default.(func() time.Time)
	// annotationnamespace.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	annotationnamespace.UpdateDefaultUpdatedAt = annotationnamespaceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// annotationnamespaceDescName is the schema descriptor for name field.
	annotationnamespaceDescName := annotationnamespaceFields[1].Descriptor()
	// annotationnamespace.NameValidator is a validator for the "name" field. It is called by the builders before save.
	annotationnamespace.NameValidator = annotationnamespaceDescName.Validators[0].(func(string) error)
	// annotationnamespaceDescPrivate is the schema descriptor for private field.
	annotationnamespaceDescPrivate := annotationnamespaceFields[3].Descriptor()
	// annotationnamespace.DefaultPrivate holds the default value on creation for the private field.
	annotationnamespace.DefaultPrivate = annotationnamespaceDescPrivate.Default.(bool)
	// annotationnamespaceDescID is the schema descriptor for id field.
	annotationnamespaceDescID := annotationnamespaceFields[0].Descriptor()
	// annotationnamespace.DefaultID holds the default value on creation for the id field.
	annotationnamespace.DefaultID = annotationnamespaceDescID.Default.(func() gidx.PrefixedID)
	metadataMixin := schema.Metadata{}.Mixin()
	metadataMixinFields0 := metadataMixin[0].Fields()
	_ = metadataMixinFields0
	metadataFields := schema.Metadata{}.Fields()
	_ = metadataFields
	// metadataDescCreatedAt is the schema descriptor for created_at field.
	metadataDescCreatedAt := metadataMixinFields0[0].Descriptor()
	// metadata.DefaultCreatedAt holds the default value on creation for the created_at field.
	metadata.DefaultCreatedAt = metadataDescCreatedAt.Default.(func() time.Time)
	// metadataDescUpdatedAt is the schema descriptor for updated_at field.
	metadataDescUpdatedAt := metadataMixinFields0[1].Descriptor()
	// metadata.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	metadata.DefaultUpdatedAt = metadataDescUpdatedAt.Default.(func() time.Time)
	// metadata.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	metadata.UpdateDefaultUpdatedAt = metadataDescUpdatedAt.UpdateDefault.(func() time.Time)
	// metadataDescID is the schema descriptor for id field.
	metadataDescID := metadataFields[0].Descriptor()
	// metadata.DefaultID holds the default value on creation for the id field.
	metadata.DefaultID = metadataDescID.Default.(func() gidx.PrefixedID)
	statusMixin := schema.Status{}.Mixin()
	statusMixinFields0 := statusMixin[0].Fields()
	_ = statusMixinFields0
	statusFields := schema.Status{}.Fields()
	_ = statusFields
	// statusDescCreatedAt is the schema descriptor for created_at field.
	statusDescCreatedAt := statusMixinFields0[0].Descriptor()
	// status.DefaultCreatedAt holds the default value on creation for the created_at field.
	status.DefaultCreatedAt = statusDescCreatedAt.Default.(func() time.Time)
	// statusDescUpdatedAt is the schema descriptor for updated_at field.
	statusDescUpdatedAt := statusMixinFields0[1].Descriptor()
	// status.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	status.DefaultUpdatedAt = statusDescUpdatedAt.Default.(func() time.Time)
	// status.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	status.UpdateDefaultUpdatedAt = statusDescUpdatedAt.UpdateDefault.(func() time.Time)
	// statusDescSource is the schema descriptor for source field.
	statusDescSource := statusFields[3].Descriptor()
	// status.SourceValidator is a validator for the "source" field. It is called by the builders before save.
	status.SourceValidator = statusDescSource.Validators[0].(func(string) error)
	// statusDescID is the schema descriptor for id field.
	statusDescID := statusFields[0].Descriptor()
	// status.DefaultID holds the default value on creation for the id field.
	status.DefaultID = statusDescID.Default.(func() gidx.PrefixedID)
	statusnamespaceMixin := schema.StatusNamespace{}.Mixin()
	statusnamespaceMixinFields0 := statusnamespaceMixin[0].Fields()
	_ = statusnamespaceMixinFields0
	statusnamespaceFields := schema.StatusNamespace{}.Fields()
	_ = statusnamespaceFields
	// statusnamespaceDescCreatedAt is the schema descriptor for created_at field.
	statusnamespaceDescCreatedAt := statusnamespaceMixinFields0[0].Descriptor()
	// statusnamespace.DefaultCreatedAt holds the default value on creation for the created_at field.
	statusnamespace.DefaultCreatedAt = statusnamespaceDescCreatedAt.Default.(func() time.Time)
	// statusnamespaceDescUpdatedAt is the schema descriptor for updated_at field.
	statusnamespaceDescUpdatedAt := statusnamespaceMixinFields0[1].Descriptor()
	// statusnamespace.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	statusnamespace.DefaultUpdatedAt = statusnamespaceDescUpdatedAt.Default.(func() time.Time)
	// statusnamespace.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	statusnamespace.UpdateDefaultUpdatedAt = statusnamespaceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// statusnamespaceDescName is the schema descriptor for name field.
	statusnamespaceDescName := statusnamespaceFields[1].Descriptor()
	// statusnamespace.NameValidator is a validator for the "name" field. It is called by the builders before save.
	statusnamespace.NameValidator = statusnamespaceDescName.Validators[0].(func(string) error)
	// statusnamespaceDescPrivate is the schema descriptor for private field.
	statusnamespaceDescPrivate := statusnamespaceFields[3].Descriptor()
	// statusnamespace.DefaultPrivate holds the default value on creation for the private field.
	statusnamespace.DefaultPrivate = statusnamespaceDescPrivate.Default.(bool)
	// statusnamespaceDescID is the schema descriptor for id field.
	statusnamespaceDescID := statusnamespaceFields[0].Descriptor()
	// statusnamespace.DefaultID holds the default value on creation for the id field.
	statusnamespace.DefaultID = statusnamespaceDescID.Default.(func() gidx.PrefixedID)
}
