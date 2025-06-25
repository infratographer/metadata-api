package graphapi_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/permissions-api/pkg/permissions/mockpermissions"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/gidx"
	"go.infratographer.com/x/testing/eventtools"

	"go.infratographer.com/metadata-api/internal/ent/generated/metadata"
	"go.infratographer.com/metadata-api/internal/ent/generated/status"
	"go.infratographer.com/metadata-api/internal/testclient"
)

var errTimeout = errors.New("timeout waiting for event")

func TestStatusUpdate(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	meta1 := MetadataBuilder{}.MustNew(ctx)
	st1 := StatusBuilder{Metadata: meta1}.MustNew(ctx)

	testCases := []struct {
		TestName    string
		NodeID      gidx.PrefixedID
		NamespaceID gidx.PrefixedID
		JSONData    json.RawMessage // optional, otherwise generated
		Source      string
		ErrorMsg    string
	}{
		{
			TestName:    "Will create status for a node we don't have metadata for",
			NodeID:      gidx.MustNewID("testing"),
			NamespaceID: StatusNamespaceBuilder{}.MustNew(ctx).ID,
			Source:      "go-tests",
		},
		{
			TestName:    "Will create status for a node that has other metadata",
			NodeID:      meta1.NodeID,
			NamespaceID: StatusNamespaceBuilder{}.MustNew(ctx).ID,
			Source:      "go-tests",
		},
		{
			TestName:    "Will create status when status already exists from a different source",
			NodeID:      meta1.NodeID,
			NamespaceID: st1.StatusNamespaceID,
			Source:      "go-tests",
		},
		{
			TestName:    "Will update status when status already exists from the same source",
			NodeID:      meta1.NodeID,
			NamespaceID: st1.StatusNamespaceID,
			Source:      st1.Source,
		},
		{
			TestName:    "Fails when namespace doesn't exist",
			NodeID:      gidx.MustNewID("testing"),
			NamespaceID: gidx.MustNewID("notreal"),
			Source:      "go-tests",
			ErrorMsg:    "status_namespace not found",
		},
		{
			TestName:    "Fails when nodeID is empty",
			NodeID:      "",
			NamespaceID: st1.StatusNamespaceID,
			Source:      "go-tests",
			ErrorMsg:    "must not be empty",
		},
		{
			TestName:    "Fails when StatusNamespaceID is empty",
			NodeID:      gidx.MustNewID("testing"),
			NamespaceID: "",
			Source:      "go-tests",
			ErrorMsg:    "must not be empty",
		},
		{
			TestName:    "Fails when source is empty",
			NodeID:      "",
			NamespaceID: st1.StatusNamespaceID,
			Source:      "",
			ErrorMsg:    "must not be empty",
		},
		{
			TestName:    "Fails when statusNamespaceID is an invalid id",
			NodeID:      gidx.MustNewID("testing"),
			NamespaceID: "a-invalid",
			Source:      "go-tests",
			ErrorMsg:    "invalid id",
		},
		{
			TestName:    "Fails when nodeID is an invalid id",
			NodeID:      "a-invalid",
			NamespaceID: st1.StatusNamespaceID,
			Source:      "go-tests",
			ErrorMsg:    "invalid id",
		},
		{
			TestName:    "Fails to update nodeID status with invalid json data",
			NodeID:      gidx.MustNewID("testing"),
			NamespaceID: st1.StatusNamespaceID,
			JSONData:    json.RawMessage(`{{}`),
			Source:      "go-tests",
			ErrorMsg:    "error calling MarshalJSON",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			if tt.JSONData == nil {
				jsonData, err := gofakeit.JSON(nil)
				tt.JSONData = json.RawMessage(jsonData)
				require.NoError(t, err)
			}

			resp, err := graphTestClient().StatusUpdate(ctx, testclient.StatusUpdateInput{NodeID: tt.NodeID, NamespaceID: tt.NamespaceID, Source: tt.Source, Data: tt.JSONData})

			if tt.ErrorMsg != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.ErrorMsg)

				return
			}

			require.NoError(t, err)
			assert.NotNil(t, resp.StatusUpdate.Status)
			assert.JSONEq(t, string(tt.JSONData), string(resp.StatusUpdate.Status.Data))

			stCount := EntClient.Status.Query().Where(status.StatusNamespaceID(tt.NamespaceID), status.Source(tt.Source), status.HasMetadataWith(metadata.NodeID(tt.NodeID))).CountX(ctx)
			assert.Equal(t, 1, stCount)
		})
	}
}

func TestStatusUpdateIsNoopWithSameValue(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	source := "go-test-noop"
	meta := MetadataBuilder{}.MustNew(ctx)
	status := StatusBuilder{
		Metadata: meta,
		Data:     json.RawMessage(`{"first": "first", "last": "last"}`),
		Source:   source,
	}.MustNew(ctx)

	jsonDataUpdate1 := json.RawMessage(`{"third":"db-sorts-by-key","first":"fun","second":"weeee"}`)
	jsonDataUpdate2 := json.RawMessage(`{"second": "weeee", "third": "db-sorts-by-key", "first": "fun"}`)
	require.JSONEq(t, string(jsonDataUpdate1), string(jsonDataUpdate2))

	// Subscribe to NATs changes
	natsCfg := NATSServer.Config.NATS
	natsCfg.QueueGroup = "event-noop-test"
	conn, err := events.NewNATSConnection(natsCfg)
	require.NoError(t, err)
	messages, err := conn.SubscribeChanges(context.Background(), ">")
	require.NoError(t, err)

	// clear out any existing events in the queue
	for {
		_, err = getSingleMessage(messages, time.Second*1)
		if err != nil {
			break
		}
	}

	// Set the status and ensure we get the data we passed back
	createdResp, err := graphTestClient().StatusUpdate(ctx, testclient.StatusUpdateInput{NodeID: meta.NodeID, NamespaceID: status.StatusNamespaceID, Source: source, Data: jsonDataUpdate1})
	require.NoError(t, err)
	assert.NotNil(t, createdResp.StatusUpdate.Status)
	assert.JSONEq(t, string(jsonDataUpdate1), string(createdResp.StatusUpdate.Status.Data))

	// Ensure event was sent with string json data and a json diff
	receivedMsg, err := getSingleMessage(messages, time.Second*1)
	require.NoError(t, err)
	require.NoError(t, receivedMsg.Error())
	require.NoError(t, receivedMsg.Ack())

	require.Equal(t, "update", receivedMsg.Message().EventType)
	require.Equal(t, eventtools.Prefix+".changes.update.status", receivedMsg.Topic())

	dataChecked := false
	for _, field := range receivedMsg.Message().FieldChanges {
		if field.Field != "data" {
			continue
		}

		assert.JSONEq(t, string(jsonDataUpdate1), field.CurrentValue)
		dataChecked = true
	}
	assert.True(t, dataChecked)

	expectedDiff := `{"first":{"changed": ["first", "fun"]}, "prop-added":{"second": "weee"}, "prop-added":{"third": "db-sorts-by-key"}, "prop-removed":{"last": "last"} }`
	assert.JSONEq(t, expectedDiff, receivedMsg.Message().AdditionalData["data-json-diff"].(string))

	// update the status, since the JSON is functionally the same, ensure we get the sorted json back
	updatedResp, err := graphTestClient().StatusUpdate(ctx, testclient.StatusUpdateInput{NodeID: meta.NodeID, NamespaceID: status.StatusNamespaceID, Source: source, Data: jsonDataUpdate2})
	require.NoError(t, err)
	assert.NotNil(t, updatedResp.StatusUpdate.Status)
	assert.JSONEq(t, string(jsonDataUpdate1), string(updatedResp.StatusUpdate.Status.Data))

	assert.NotEqual(t, createdResp.StatusUpdate.Status.UpdatedAt, updatedResp.StatusUpdate.Status.UpdatedAt)

	// Ensure no eventa was sent for update event since there was no update
	receivedMsg, err = getSingleMessage(messages, time.Second*1)
	assert.Nil(t, receivedMsg)
	assert.ErrorIs(t, err, errTimeout)
}

func TestStatusDelete(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	meta1 := MetadataBuilder{}.MustNew(ctx)
	st1 := StatusBuilder{Metadata: meta1}.MustNew(ctx)
	st2 := StatusBuilder{Metadata: meta1}.MustNew(ctx)

	testCases := []struct {
		TestName    string
		NodeID      gidx.PrefixedID
		NamespaceID gidx.PrefixedID
		Source      string
		ErrorMsg    string
	}{
		{
			TestName:    "Will delete status when found",
			NodeID:      meta1.NodeID,
			NamespaceID: st1.StatusNamespaceID,
			Source:      st1.Source,
		},
		{
			TestName:    "Will return an error if the status doesn't exists for the given source and namespace",
			NodeID:      meta1.NodeID,
			NamespaceID: st2.StatusNamespaceID,
			Source:      "this-is-not-source-you-are-looking-for",
			ErrorMsg:    "status not found",
		},
		{
			TestName:    "fails when NodeID is empty",
			NodeID:      "",
			NamespaceID: st2.StatusNamespaceID,
			Source:      "unit-test",
			ErrorMsg:    "must not be empty",
		},
		{
			TestName:    "fails when NamespaceID is empty",
			NodeID:      meta1.NodeID,
			NamespaceID: "",
			Source:      "unit-test",
			ErrorMsg:    "must not be empty",
		},
		{
			TestName:    "fails when NodeID is an invalid gidx",
			NodeID:      "a-testing",
			NamespaceID: st2.StatusNamespaceID,
			Source:      "unit-test",
			ErrorMsg:    "invalid id",
		},
		{
			TestName:    "fails when StatusNamespaceID is an invalid gidx",
			NodeID:      gidx.MustNewID("testing"),
			NamespaceID: "a-testing",
			Source:      "unit-test",
			ErrorMsg:    "invalid id",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().StatusDelete(ctx, testclient.StatusDeleteInput{NodeID: tt.NodeID, NamespaceID: tt.NamespaceID, Source: tt.Source})

			if tt.ErrorMsg != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.ErrorMsg)

				return
			}

			require.NoError(t, err)
			assert.NotNil(t, resp.StatusDelete)
			assert.NotNil(t, resp.StatusDelete.DeletedID)

			count := EntClient.Status.Query().Where(status.Source(tt.Source), status.HasMetadataWith(metadata.NodeID(tt.NodeID))).CountX(ctx)
			assert.Equal(t, 0, count)
		})
	}
}

func getSingleMessage[T any](messages <-chan events.Message[T], timeout time.Duration) (events.Message[T], error) {
	select {
	case message := <-messages:
		return message, nil
	case <-time.After(timeout):
		return nil, errTimeout
	}
}
