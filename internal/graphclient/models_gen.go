// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphclient

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"go.infratographer.com/x/gidx"
)

// An object with an ID.
// Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
type Node interface {
	IsNode()
	// The id of the object.
	GetID() gidx.PrefixedID
}

type Entity interface {
	IsEntity()
}

type Annotation struct {
	// ID for the annotation.
	ID        gidx.PrefixedID `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	// ID of the metadata of this annotation
	MetadataID gidx.PrefixedID `json:"metadataID"`
	// JSON formatted data of this annotation.
	Data      json.RawMessage     `json:"data"`
	Namespace AnnotationNamespace `json:"namespace"`
	Metadata  Metadata            `json:"metadata"`
}

func (Annotation) IsNode() {}

// The id of the object.
func (this Annotation) GetID() gidx.PrefixedID { return this.ID }

func (Annotation) IsEntity() {}

// A connection to a list of items.
type AnnotationConnection struct {
	// A list of edges.
	Edges []*AnnotationEdge `json:"edges,omitempty"`
	// Information to aid in pagination.
	PageInfo PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// Input information to delete an annotation.
type AnnotationDeleteInput struct {
	// The node ID for this annotation.
	NodeID gidx.PrefixedID `json:"nodeID"`
	// The namespace ID for this annotation.
	NamespaceID gidx.PrefixedID `json:"namespaceID"`
}

// Return response from annotationDelete
type AnnotationDeleteResponse struct {
	// The ID of the unset annotation.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// An edge in a connection.
type AnnotationEdge struct {
	// The item at the end of the edge.
	Node *Annotation `json:"node,omitempty"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

type AnnotationNamespace struct {
	// The ID for the annotation namespace.
	ID        gidx.PrefixedID `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	// The name of the annotation namespace.
	Name string `json:"name"`
	// Flag for if this namespace is private.
	Private     bool          `json:"private"`
	Annotations []*Annotation `json:"annotations,omitempty"`
	// The tenant of the annotation namespace.
	Tenant Tenant `json:"tenant"`
}

func (AnnotationNamespace) IsNode() {}

// The id of the object.
func (this AnnotationNamespace) GetID() gidx.PrefixedID { return this.ID }

func (AnnotationNamespace) IsEntity() {}

// A connection to a list of items.
type AnnotationNamespaceConnection struct {
	// A list of edges.
	Edges []*AnnotationNamespaceEdge `json:"edges,omitempty"`
	// Information to aid in pagination.
	PageInfo PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// Return response from annotationNamespaceCreate
type AnnotationNamespaceCreatePayload struct {
	// The created annotation namespace.
	AnnotationNamespace AnnotationNamespace `json:"annotationNamespace"`
}

// Return response from annotationNamespaceDelete
type AnnotationNamespaceDeletePayload struct {
	// The ID of the deleted annotation namespace.
	DeletedID gidx.PrefixedID `json:"deletedID"`
	// The count of annotations deleted
	AnnotationDeletedCount int64 `json:"annotationDeletedCount"`
}

// An edge in a connection.
type AnnotationNamespaceEdge struct {
	// The item at the end of the edge.
	Node *AnnotationNamespace `json:"node,omitempty"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

// Ordering options for AnnotationNamespace connections
type AnnotationNamespaceOrder struct {
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
	// The field by which to order AnnotationNamespaces.
	Field AnnotationNamespaceOrderField `json:"field"`
}

// Return response from annotationNamespaceUpdate
type AnnotationNamespaceUpdatePayload struct {
	// The updated annotation namespace.
	AnnotationNamespace AnnotationNamespace `json:"annotationNamespace"`
}

// AnnotationNamespaceWhereInput is used for filtering AnnotationNamespace objects.
// Input was generated by ent.
type AnnotationNamespaceWhereInput struct {
	Not *AnnotationNamespaceWhereInput   `json:"not,omitempty"`
	And []*AnnotationNamespaceWhereInput `json:"and,omitempty"`
	Or  []*AnnotationNamespaceWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNeq   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGt    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGte   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLt    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLte   *gidx.PrefixedID  `json:"idLTE,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// updated_at field predicates
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq   *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt    *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte   *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt    *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte   *time.Time   `json:"updatedAtLTE,omitempty"`
	// name field predicates
	Name             *string  `json:"name,omitempty"`
	NameNeq          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGt           *string  `json:"nameGT,omitempty"`
	NameGte          *string  `json:"nameGTE,omitempty"`
	NameLt           *string  `json:"nameLT,omitempty"`
	NameLte          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
	// annotations edge predicates
	HasAnnotations     *bool                   `json:"hasAnnotations,omitempty"`
	HasAnnotationsWith []*AnnotationWhereInput `json:"hasAnnotationsWith,omitempty"`
}

// Ordering options for Annotation connections
type AnnotationOrder struct {
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
	// The field by which to order Annotations.
	Field AnnotationOrderField `json:"field"`
}

// Input information to update an annotation.
type AnnotationUpdateInput struct {
	// The node ID for this annotation.
	NodeID gidx.PrefixedID `json:"nodeID"`
	// The namespace ID for this annotation.
	NamespaceID gidx.PrefixedID `json:"namespaceID"`
	// The data to save in this annotation.
	Data json.RawMessage `json:"data"`
}

// Return response from annotationUpdate
type AnnotationUpdateResponse struct {
	// The set annotation.
	Annotation Annotation `json:"annotation"`
}

// AnnotationWhereInput is used for filtering Annotation objects.
// Input was generated by ent.
type AnnotationWhereInput struct {
	Not *AnnotationWhereInput   `json:"not,omitempty"`
	And []*AnnotationWhereInput `json:"and,omitempty"`
	Or  []*AnnotationWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNeq   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGt    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGte   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLt    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLte   *gidx.PrefixedID  `json:"idLTE,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// updated_at field predicates
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq   *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt    *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte   *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt    *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte   *time.Time   `json:"updatedAtLTE,omitempty"`
	// namespace edge predicates
	HasNamespace     *bool                            `json:"hasNamespace,omitempty"`
	HasNamespaceWith []*AnnotationNamespaceWhereInput `json:"hasNamespaceWith,omitempty"`
	// metadata edge predicates
	HasMetadata     *bool                 `json:"hasMetadata,omitempty"`
	HasMetadataWith []*MetadataWhereInput `json:"hasMetadataWith,omitempty"`
}

// Input information to create an annotation namespace.
type CreateAnnotationNamespaceInput struct {
	// The name of the annotation namespace.
	Name string `json:"name"`
	// The ID for the tenant for this annotation namespace.
	TenantID gidx.PrefixedID `json:"tenantID"`
	// Flag for if this namespace is private.
	Private *bool `json:"private,omitempty"`
}

// Input information to create a status namespace.
type CreateStatusInput struct {
	Source string `json:"source"`
	// JSON formatted data of this annotation.
	Data        json.RawMessage `json:"data"`
	NamespaceID gidx.PrefixedID `json:"namespaceID"`
	MetadataID  gidx.PrefixedID `json:"metadataID"`
}

// Input information to create a status namespace.
type CreateStatusNamespaceInput struct {
	// The name of the status namespace.
	Name string `json:"name"`
	// The ID for the tenant for this status namespace.
	ResourceProviderID gidx.PrefixedID `json:"resourceProviderID"`
	// Flag for if this namespace is private.
	Private *bool `json:"private,omitempty"`
}

type Entity0 struct {
	FindAnnotationByID          Annotation          `json:"findAnnotationByID"`
	FindAnnotationNamespaceByID AnnotationNamespace `json:"findAnnotationNamespaceByID"`
	FindMetadataByID            Metadata            `json:"findMetadataByID"`
	FindResourceProviderByID    ResourceProvider    `json:"findResourceProviderByID"`
	FindStatusByID              Status              `json:"findStatusByID"`
	FindStatusNamespaceByID     StatusNamespace     `json:"findStatusNamespaceByID"`
	FindTenantByID              Tenant              `json:"findTenantByID"`
}

type Metadata struct {
	// ID for the metadata.
	ID        gidx.PrefixedID `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	// ID of the node for this metadata
	NodeID      gidx.PrefixedID `json:"nodeID"`
	Annotations []*Annotation   `json:"annotations,omitempty"`
	Statuses    []*Status       `json:"statuses,omitempty"`
}

func (Metadata) IsNode() {}

// The id of the object.
func (this Metadata) GetID() gidx.PrefixedID { return this.ID }

func (Metadata) IsEntity() {}

// A connection to a list of items.
type MetadataConnection struct {
	// A list of edges.
	Edges []*MetadataEdge `json:"edges,omitempty"`
	// Information to aid in pagination.
	PageInfo PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// An edge in a connection.
type MetadataEdge struct {
	// The item at the end of the edge.
	Node *Metadata `json:"node,omitempty"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

// Ordering options for Metadata connections
type MetadataOrder struct {
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
	// The field by which to order MetadataSlice.
	Field MetadataOrderField `json:"field"`
}

// MetadataWhereInput is used for filtering Metadata objects.
// Input was generated by ent.
type MetadataWhereInput struct {
	Not *MetadataWhereInput   `json:"not,omitempty"`
	And []*MetadataWhereInput `json:"and,omitempty"`
	Or  []*MetadataWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNeq   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGt    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGte   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLt    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLte   *gidx.PrefixedID  `json:"idLTE,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// updated_at field predicates
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq   *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt    *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte   *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt    *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte   *time.Time   `json:"updatedAtLTE,omitempty"`
	// annotations edge predicates
	HasAnnotations     *bool                   `json:"hasAnnotations,omitempty"`
	HasAnnotationsWith []*AnnotationWhereInput `json:"hasAnnotationsWith,omitempty"`
	// statuses edge predicates
	HasStatuses     *bool               `json:"hasStatuses,omitempty"`
	HasStatusesWith []*StatusWhereInput `json:"hasStatusesWith,omitempty"`
}

// Information about pagination in a connection.
// https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor *string `json:"startCursor,omitempty"`
	// When paginating forwards, the cursor to continue.
	EndCursor *string `json:"endCursor,omitempty"`
}

type ResourceProvider struct {
	ID               gidx.PrefixedID           `json:"id"`
	StatusNamespaces StatusNamespaceConnection `json:"statusNamespaces"`
}

func (ResourceProvider) IsEntity() {}

type Status struct {
	ID        gidx.PrefixedID `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	// ID of the metadata of this status
	MetadataID        gidx.PrefixedID `json:"metadataID"`
	StatusNamespaceID gidx.PrefixedID `json:"statusNamespaceID"`
	Source            string          `json:"source"`
	// JSON formatted data of this annotation.
	Data      json.RawMessage `json:"data"`
	Namespace StatusNamespace `json:"namespace"`
	Metadata  Metadata        `json:"metadata"`
}

func (Status) IsNode() {}

// The id of the object.
func (this Status) GetID() gidx.PrefixedID { return this.ID }

func (Status) IsEntity() {}

// A connection to a list of items.
type StatusConnection struct {
	// A list of edges.
	Edges []*StatusEdge `json:"edges,omitempty"`
	// Information to aid in pagination.
	PageInfo PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// Input information to delete an status.
type StatusDeleteInput struct {
	// The node ID for this status.
	NodeID gidx.PrefixedID `json:"nodeID"`
	// The namespace ID for this status.
	NamespaceID gidx.PrefixedID `json:"namespaceID"`
	// The source for this status.
	Source string `json:"source"`
}

// Return response from statusDelete
type StatusDeleteResponse struct {
	// The ID of the unset status.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// An edge in a connection.
type StatusEdge struct {
	// The item at the end of the edge.
	Node *Status `json:"node,omitempty"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

type StatusNamespace struct {
	// The ID for the status namespace.
	ID        gidx.PrefixedID `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	// The name of the status namespace.
	Name string `json:"name"`
	// Flag for if this namespace is private.
	Private bool `json:"private"`
	// The resource provider of the status namespace.
	ResourceProvider ResourceProvider `json:"resourceProvider"`
	// The tenant of the annotation namespace.
	Tenant Tenant `json:"tenant"`
}

func (StatusNamespace) IsNode() {}

// The id of the object.
func (this StatusNamespace) GetID() gidx.PrefixedID { return this.ID }

func (StatusNamespace) IsEntity() {}

// A connection to a list of items.
type StatusNamespaceConnection struct {
	// A list of edges.
	Edges []*StatusNamespaceEdge `json:"edges,omitempty"`
	// Information to aid in pagination.
	PageInfo PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// Return response from statusNamespaceCreate
type StatusNamespaceCreatePayload struct {
	// The created status namespace.
	StatusNamespace StatusNamespace `json:"statusNamespace"`
}

// Return response from statusNamespaceDelete
type StatusNamespaceDeletePayload struct {
	// The ID of the deleted status namespace.
	DeletedID gidx.PrefixedID `json:"deletedID"`
	// The count of statuss deleted
	StatusDeletedCount int64 `json:"statusDeletedCount"`
}

// An edge in a connection.
type StatusNamespaceEdge struct {
	// The item at the end of the edge.
	Node *StatusNamespace `json:"node,omitempty"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

// Ordering options for StatusNamespace connections
type StatusNamespaceOrder struct {
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
	// The field by which to order StatusNamespaces.
	Field StatusNamespaceOrderField `json:"field"`
}

// Return response from statusNamespaceUpdate
type StatusNamespaceUpdatePayload struct {
	// The updated status namespace.
	StatusNamespace StatusNamespace `json:"statusNamespace"`
}

// StatusNamespaceWhereInput is used for filtering StatusNamespace objects.
// Input was generated by ent.
type StatusNamespaceWhereInput struct {
	Not *StatusNamespaceWhereInput   `json:"not,omitempty"`
	And []*StatusNamespaceWhereInput `json:"and,omitempty"`
	Or  []*StatusNamespaceWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNeq   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGt    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGte   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLt    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLte   *gidx.PrefixedID  `json:"idLTE,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// updated_at field predicates
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq   *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt    *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte   *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt    *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte   *time.Time   `json:"updatedAtLTE,omitempty"`
	// name field predicates
	Name             *string  `json:"name,omitempty"`
	NameNeq          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGt           *string  `json:"nameGT,omitempty"`
	NameGte          *string  `json:"nameGTE,omitempty"`
	NameLt           *string  `json:"nameLT,omitempty"`
	NameLte          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
}

// Ordering options for Status connections
type StatusOrder struct {
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
	// The field by which to order StatusSlice.
	Field StatusOrderField `json:"field"`
}

// Input information to update an status.
type StatusUpdateInput struct {
	// The node ID for this status.
	NodeID gidx.PrefixedID `json:"nodeID"`
	// The namespace ID for this status.
	NamespaceID gidx.PrefixedID `json:"namespaceID"`
	// The source for this status.
	Source string `json:"source"`
	// The data to save in this status.
	Data json.RawMessage `json:"data"`
}

// Return response from statusUpdate
type StatusUpdateResponse struct {
	// The set status.
	Status Status `json:"status"`
}

// StatusWhereInput is used for filtering Status objects.
// Input was generated by ent.
type StatusWhereInput struct {
	Not *StatusWhereInput   `json:"not,omitempty"`
	And []*StatusWhereInput `json:"and,omitempty"`
	Or  []*StatusWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNeq   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGt    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGte   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLt    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLte   *gidx.PrefixedID  `json:"idLTE,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// updated_at field predicates
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq   *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt    *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte   *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt    *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte   *time.Time   `json:"updatedAtLTE,omitempty"`
	// source field predicates
	Source             *string  `json:"source,omitempty"`
	SourceNeq          *string  `json:"sourceNEQ,omitempty"`
	SourceIn           []string `json:"sourceIn,omitempty"`
	SourceNotIn        []string `json:"sourceNotIn,omitempty"`
	SourceGt           *string  `json:"sourceGT,omitempty"`
	SourceGte          *string  `json:"sourceGTE,omitempty"`
	SourceLt           *string  `json:"sourceLT,omitempty"`
	SourceLte          *string  `json:"sourceLTE,omitempty"`
	SourceContains     *string  `json:"sourceContains,omitempty"`
	SourceHasPrefix    *string  `json:"sourceHasPrefix,omitempty"`
	SourceHasSuffix    *string  `json:"sourceHasSuffix,omitempty"`
	SourceEqualFold    *string  `json:"sourceEqualFold,omitempty"`
	SourceContainsFold *string  `json:"sourceContainsFold,omitempty"`
	// namespace edge predicates
	HasNamespace     *bool                        `json:"hasNamespace,omitempty"`
	HasNamespaceWith []*StatusNamespaceWhereInput `json:"hasNamespaceWith,omitempty"`
	// metadata edge predicates
	HasMetadata     *bool                 `json:"hasMetadata,omitempty"`
	HasMetadataWith []*MetadataWhereInput `json:"hasMetadataWith,omitempty"`
}

type Tenant struct {
	ID                   gidx.PrefixedID               `json:"id"`
	AnnotationNamespaces AnnotationNamespaceConnection `json:"annotationNamespaces"`
	StatusNamespaces     StatusNamespaceConnection     `json:"statusNamespaces"`
}

func (Tenant) IsEntity() {}

// Input information to update an annotation namespace.
type UpdateAnnotationNamespaceInput struct {
	// The name of the annotation namespace.
	Name *string `json:"name,omitempty"`
	// Flag for if this namespace is private.
	Private *bool `json:"private,omitempty"`
}

// Input information to update a status namespace.
type UpdateStatusInput struct {
	// JSON formatted data of this annotation.
	Data       json.RawMessage `json:"data,omitempty"`
	AppendData json.RawMessage `json:"appendData,omitempty"`
}

// Input information to update a status namespace.
type UpdateStatusNamespaceInput struct {
	// The name of the status namespace.
	Name *string `json:"name,omitempty"`
	// Flag for if this namespace is private.
	Private *bool `json:"private,omitempty"`
}

type Service struct {
	Sdl *string `json:"sdl,omitempty"`
}

// Properties by which AnnotationNamespace connections can be ordered.
type AnnotationNamespaceOrderField string

const (
	AnnotationNamespaceOrderFieldID        AnnotationNamespaceOrderField = "ID"
	AnnotationNamespaceOrderFieldCreatedAt AnnotationNamespaceOrderField = "CREATED_AT"
	AnnotationNamespaceOrderFieldUpdatedAt AnnotationNamespaceOrderField = "UPDATED_AT"
	AnnotationNamespaceOrderFieldName      AnnotationNamespaceOrderField = "NAME"
	AnnotationNamespaceOrderFieldTenant    AnnotationNamespaceOrderField = "TENANT"
	AnnotationNamespaceOrderFieldPrivate   AnnotationNamespaceOrderField = "PRIVATE"
)

var AllAnnotationNamespaceOrderField = []AnnotationNamespaceOrderField{
	AnnotationNamespaceOrderFieldID,
	AnnotationNamespaceOrderFieldCreatedAt,
	AnnotationNamespaceOrderFieldUpdatedAt,
	AnnotationNamespaceOrderFieldName,
	AnnotationNamespaceOrderFieldTenant,
	AnnotationNamespaceOrderFieldPrivate,
}

func (e AnnotationNamespaceOrderField) IsValid() bool {
	switch e {
	case AnnotationNamespaceOrderFieldID, AnnotationNamespaceOrderFieldCreatedAt, AnnotationNamespaceOrderFieldUpdatedAt, AnnotationNamespaceOrderFieldName, AnnotationNamespaceOrderFieldTenant, AnnotationNamespaceOrderFieldPrivate:
		return true
	}
	return false
}

func (e AnnotationNamespaceOrderField) String() string {
	return string(e)
}

func (e *AnnotationNamespaceOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AnnotationNamespaceOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AnnotationNamespaceOrderField", str)
	}
	return nil
}

func (e AnnotationNamespaceOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which Annotation connections can be ordered.
type AnnotationOrderField string

const (
	AnnotationOrderFieldCreatedAt AnnotationOrderField = "CREATED_AT"
	AnnotationOrderFieldUpdatedAt AnnotationOrderField = "UPDATED_AT"
)

var AllAnnotationOrderField = []AnnotationOrderField{
	AnnotationOrderFieldCreatedAt,
	AnnotationOrderFieldUpdatedAt,
}

func (e AnnotationOrderField) IsValid() bool {
	switch e {
	case AnnotationOrderFieldCreatedAt, AnnotationOrderFieldUpdatedAt:
		return true
	}
	return false
}

func (e AnnotationOrderField) String() string {
	return string(e)
}

func (e *AnnotationOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AnnotationOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AnnotationOrderField", str)
	}
	return nil
}

func (e AnnotationOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which Metadata connections can be ordered.
type MetadataOrderField string

const (
	MetadataOrderFieldCreatedAt MetadataOrderField = "CREATED_AT"
	MetadataOrderFieldUpdatedAt MetadataOrderField = "UPDATED_AT"
)

var AllMetadataOrderField = []MetadataOrderField{
	MetadataOrderFieldCreatedAt,
	MetadataOrderFieldUpdatedAt,
}

func (e MetadataOrderField) IsValid() bool {
	switch e {
	case MetadataOrderFieldCreatedAt, MetadataOrderFieldUpdatedAt:
		return true
	}
	return false
}

func (e MetadataOrderField) String() string {
	return string(e)
}

func (e *MetadataOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MetadataOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MetadataOrderField", str)
	}
	return nil
}

func (e MetadataOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Possible directions in which to order a list of items when provided an `orderBy` argument.
type OrderDirection string

const (
	// Specifies an ascending order for a given `orderBy` argument.
	OrderDirectionAsc OrderDirection = "ASC"
	// Specifies a descending order for a given `orderBy` argument.
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which StatusNamespace connections can be ordered.
type StatusNamespaceOrderField string

const (
	StatusNamespaceOrderFieldID        StatusNamespaceOrderField = "ID"
	StatusNamespaceOrderFieldCreatedAt StatusNamespaceOrderField = "CREATED_AT"
	StatusNamespaceOrderFieldUpdatedAt StatusNamespaceOrderField = "UPDATED_AT"
	StatusNamespaceOrderFieldName      StatusNamespaceOrderField = "NAME"
	StatusNamespaceOrderFieldTenant    StatusNamespaceOrderField = "TENANT"
	StatusNamespaceOrderFieldPrivate   StatusNamespaceOrderField = "PRIVATE"
)

var AllStatusNamespaceOrderField = []StatusNamespaceOrderField{
	StatusNamespaceOrderFieldID,
	StatusNamespaceOrderFieldCreatedAt,
	StatusNamespaceOrderFieldUpdatedAt,
	StatusNamespaceOrderFieldName,
	StatusNamespaceOrderFieldTenant,
	StatusNamespaceOrderFieldPrivate,
}

func (e StatusNamespaceOrderField) IsValid() bool {
	switch e {
	case StatusNamespaceOrderFieldID, StatusNamespaceOrderFieldCreatedAt, StatusNamespaceOrderFieldUpdatedAt, StatusNamespaceOrderFieldName, StatusNamespaceOrderFieldTenant, StatusNamespaceOrderFieldPrivate:
		return true
	}
	return false
}

func (e StatusNamespaceOrderField) String() string {
	return string(e)
}

func (e *StatusNamespaceOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StatusNamespaceOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StatusNamespaceOrderField", str)
	}
	return nil
}

func (e StatusNamespaceOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which Status connections can be ordered.
type StatusOrderField string

const (
	StatusOrderFieldCreatedAt StatusOrderField = "CREATED_AT"
	StatusOrderFieldUpdatedAt StatusOrderField = "UPDATED_AT"
)

var AllStatusOrderField = []StatusOrderField{
	StatusOrderFieldCreatedAt,
	StatusOrderFieldUpdatedAt,
}

func (e StatusOrderField) IsValid() bool {
	switch e {
	case StatusOrderFieldCreatedAt, StatusOrderFieldUpdatedAt:
		return true
	}
	return false
}

func (e StatusOrderField) String() string {
	return string(e)
}

func (e *StatusOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StatusOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StatusOrderField", str)
	}
	return nil
}

func (e StatusOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
