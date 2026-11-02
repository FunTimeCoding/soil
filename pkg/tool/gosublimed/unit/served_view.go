package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"testing"
)

func servedView(
	t *testing.T,
	content string,
) *client.ClientWithResponses {
	t.Helper()
	result := restClient(t)
	_, e := result.CreateViewWithResponse(
		context.Background(),
		client.CreateViewJSONRequestBody{Title: "notes", Content: content},
	)
	assert.FatalOnError(t, e)

	return result
}
