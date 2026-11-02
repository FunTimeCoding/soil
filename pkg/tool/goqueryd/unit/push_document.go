package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"net/http"
	"testing"
)

func pushDocument(
	t *testing.T,
	c *client.ClientWithResponses,
	collection string,
	path string,
	body string,
) {
	t.Helper()
	answer, e := c.PostDocumentWithResponse(
		t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: collection,
			Path:       path,
			Body:       body,
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, answer.StatusCode())
}
