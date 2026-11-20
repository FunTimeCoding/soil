package document_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"net/http"
)

func (o *Tester) Push(
	c *client.ClientWithResponses,
	collection string,
	path string,
	body string,
) {
	o.t.Helper()
	answer, e := c.PostDocumentWithResponse(
		o.t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: collection,
			Path:       path,
			Body:       body,
		},
	)
	assert.FatalOnError(o.t, e)
	assert.Integer(o.t, http.StatusOK, answer.StatusCode())
}
