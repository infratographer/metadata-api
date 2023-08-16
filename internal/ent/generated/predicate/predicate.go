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

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// Annotation is the predicate function for annotation builders.
type Annotation func(*sql.Selector)

// AnnotationNamespace is the predicate function for annotationnamespace builders.
type AnnotationNamespace func(*sql.Selector)

// Metadata is the predicate function for metadata builders.
type Metadata func(*sql.Selector)

// Status is the predicate function for status builders.
type Status func(*sql.Selector)

// StatusNamespace is the predicate function for statusnamespace builders.
type StatusNamespace func(*sql.Selector)
