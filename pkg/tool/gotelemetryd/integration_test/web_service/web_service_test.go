//go:build local

package web_service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/integration_test/web_service_tester"
	"net/http"
	"testing"
)

func TestWebService(t *testing.T) {
	o := web_service_tester.New(t)
	defer o.Close()
	c := o.Client
	x := context.Background()
	post, e := c.PostEventWithResponse(
		x,
		client.PostEventJSONRequestBody{
			Tool:    "save_memory",
			Surface: "model_context",
			Actor:   "Reva",
			Outcome: "success",
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, post.StatusCode())
	assert.NotNil(t, post.JSON200)
	_, e = c.PostEventWithResponse(
		x,
		client.PostEventJSONRequestBody{
			Tool:    "fleet_deploy",
			Surface: "command_line",
			Actor:   "Hana",
			Outcome: "success",
		},
	)
	assert.FatalOnError(t, e)
	events, e := c.GetEventsWithResponse(
		x,
		&client.GetEventsParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, events.StatusCode())
	assert.Count(t, 2, *events.JSON200)
	filtered, e := c.GetEventsWithResponse(
		x,
		&client.GetEventsParams{Tool: new("save_memory")},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *filtered.JSON200)
	summary, e := c.GetSummaryWithResponse(
		x,
		&client.GetSummaryParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, summary.StatusCode())
	assert.NotEmpty(t, *summary.JSON200)
}
