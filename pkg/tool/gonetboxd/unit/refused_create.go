package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/refusing_source"
	"testing"
)

func refusedCreate(
	t *testing.T,
	refusal error,
) (int, int) {
	t.Helper()
	r := memory.New()
	c := newClient(t, newNetboxServer(t, refusing_source.New(refusal), r))
	response, e := c.CreateSiteWithResponse(
		context.Background(),
		client.CreateSiteJSONRequestBody{Name: "alfa"},
	)
	assert.FatalOnError(t, e)

	return response.StatusCode(), len(r.Events())
}
