package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"testing"
)

func newMock(t *testing.T) *mock_client.Client {
	t.Helper()
	result := mock_client.New()
	t.Cleanup(result.Close)
	_, e := result.PostDocument(
		t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: "alfa",
			Path:       "kilo.md",
			Body:       "lima",
		},
	)
	assert.FatalOnError(t, e)
	result.ContextsAdded = append(
		result.ContextsAdded,
		client.PostContextJSONRequestBody{
			Collection: "alfa",
			PathPrefix: "kilo",
		},
	)

	return result
}
