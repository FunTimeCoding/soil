package integration_test

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/integration_test/web_service_tester"
	"net/http"
	"testing"
)

func TestWebService(t *testing.T) {
	o := web_service_tester.New(t)
	defer o.Close()
	c := o.Client
	x := context.Background()
	system := "worker1"
	service := "nginx"
	description := "nginx was unresponsive"
	create, e := c.PostEntryWithResponse(
		x,
		client.PostEntryJSONRequestBody{
			Action:      "restarted web server",
			User:        "jdoe",
			System:      &system,
			Service:     &service,
			Description: &description,
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, create.StatusCode())
	assert.NotNil(t, create.JSON200)
	assert.String(t, "restarted web server", create.JSON200.Action)
	assert.String(t, "jdoe", create.JSON200.User)
	identifier := create.JSON200.Id
	status, e := c.GetStatusWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, status.StatusCode())
	assert.Integer(t, 1, status.JSON200.TotalEntries)
	entries, e := c.GetEntriesWithResponse(x, &client.GetEntriesParams{})
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, entries.StatusCode())
	assert.Count(t, 1, *entries.JSON200)
	newAction := "cleared and documented"
	update, e := c.UpdateEntryWithResponse(
		x,
		identifier,
		client.UpdateEntryJSONRequestBody{
			Action:      newAction,
			User:        "jdoe",
			System:      &system,
			Service:     &service,
			Description: &description,
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, update.StatusCode())
	assert.String(t, newAction, update.JSON200.Action)
	entries, e = c.GetEntriesWithResponse(x, &client.GetEntriesParams{})
	assert.FatalOnError(t, e)
	assert.String(t, newAction, (*entries.JSON200)[0].Action)
	del, e := c.DeleteEntryWithResponse(x, identifier)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNoContent, del.StatusCode())
	status, e = c.GetStatusWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, status.JSON200.TotalEntries)
}
