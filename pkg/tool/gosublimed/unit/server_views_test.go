package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"net/http"
	"testing"
)

func TestViewLifecycle(t *testing.T) {
	c := restClient(t)
	created, e := c.CreateViewWithResponse(
		context.Background(),
		client.CreateViewJSONRequestBody{Title: "notes", Content: "alfa"},
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, created.JSON200)
	assert.Integer(t, 1, created.JSON200.ViewId)
	assert.String(t, "notes", created.JSON200.Title)
	list, e := c.GetViewsWithResponse(context.Background())
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, list.JSON200)
	assert.Count(t, 1, *list.JSON200)
	edited, e := c.EditViewWithResponse(
		context.Background(),
		1,
		client.EditViewJSONRequestBody{OldString: "alfa", NewString: "bravo"},
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, edited.JSON200)
	assert.FatalNotNil(t, edited.JSON200.Text)
	assert.String(t, "bravo", *edited.JSON200.Text)
	read, e := c.GetViewWithResponse(context.Background(), 1)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, read.JSON200)
	assert.Integer(t, 1, read.JSON200.ViewId)
	closed, e := c.CloseViewWithResponse(context.Background(), 1)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNoContent, closed.StatusCode())
	list, e = c.GetViewsWithResponse(context.Background())
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, list.JSON200)
	assert.Count(t, 0, *list.JSON200)
}
