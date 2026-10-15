package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/mock_client"
	"net/http"
	"testing"
)

func TestConfigurationContextRoundTrip(t *testing.T) {
	x := context.Background()
	r := memory.New()
	c := newClient(t, newNetboxServer(t, mock_client.New(), r))
	created, e := c.CreateConfigurationContextWithResponse(
		x,
		client.CreateConfigurationContextRequest{
			Name:    "alfa",
			Payload: map[string]any{"bravo": "charlie"},
			Assignment: &client.ConfigurationAssignment{
				Roles: &[]string{"golf"},
			},
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusCreated, created.StatusCode())
	assert.Integer(t, 1000, created.JSON201.Weight)
	updated, e := c.UpdateConfigurationContextWithResponse(
		x,
		"alfa",
		client.UpdateConfigurationContextRequest{
			Payload: &map[string]any{"delta": "echo"},
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, updated.StatusCode())
	assert.Any(t, map[string]any{"delta": "echo"}, updated.JSON200.Payload)
	assert.Any(t, []string{"golf"}, *updated.JSON200.Assignment.Roles)
	listed, e := c.ListConfigurationContextsWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *listed.JSON200)
	deleted, e := c.DeleteConfigurationContextWithResponse(x, "alfa")
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNoContent, deleted.StatusCode())
	gone, e := c.GetConfigurationContextWithResponse(x, "alfa")
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, gone.StatusCode())
	assert.Count(t, 0, r.Events())
}

func TestUnknownConfigurationContextIsNotFoundWithoutSentry(t *testing.T) {
	r := memory.New()
	c := newClient(t, newNetboxServer(t, mock_client.New(), r))
	response, e := c.UpdateConfigurationContextWithResponse(
		context.Background(),
		"alfa",
		client.UpdateConfigurationContextRequest{Weight: new(int32(500))},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.StringContains(t, "alfa", response.JSON404.Error)
	assert.Count(t, 0, r.Events())
}

func TestConfigurationContextToolsKeepAssignmentOnPayloadUpdate(t *testing.T) {
	m := newNetboxServer(t, mock_client.New(), memory.New()).Connect(t)
	defer m.Close()
	m.MustCallToolNotation(
		constant.CreateConfigurationContext,
		map[string]any{
			"name":           "alfa",
			constant.Payload: map[string]any{"bravo": "charlie"},
			constant.Roles:   []any{"golf"},
		},
	)
	m.MustCallToolNotation(
		constant.UpdateConfigurationContext,
		map[string]any{
			"name":           "alfa",
			constant.Payload: map[string]any{"delta": "echo"},
		},
	)
	result := m.MustCallToolNotation(
		constant.GetConfigurationContext,
		map[string]any{"name": "alfa"},
	)
	assert.Any(t, map[string]any{"delta": "echo"}, result[constant.Payload])
	assert.Any(
		t,
		[]any{"golf"},
		result["assignment"].(map[string]any)[constant.Roles],
	)
}

func TestConfigurationContextToolRefusesMissingPayload(t *testing.T) {
	m := newNetboxServer(t, mock_client.New(), memory.New()).Connect(t)
	defer m.Close()
	text, e := m.CallToolError(
		constant.CreateConfigurationContext,
		map[string]any{"name": "alfa"},
	)

	if e != nil {
		text = e.Error()
	}

	assert.StringContains(t, "payload", text)
}
