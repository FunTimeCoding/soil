package client_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/unit/base"
	"testing"
)

func ServedView(
	t *testing.T,
	content string,
) *client.ClientWithResponses {
	t.Helper()
	result := base.Rest(t, base.New(t))
	_, e := result.CreateViewWithResponse(
		context.Background(),
		client.CreateViewJSONRequestBody{Title: "notes", Content: content},
	)
	assert.FatalOnError(t, e)

	return result
}
