package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/held_store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"net/http"
	"testing"
)

func TestPushWhileAnotherWriterHoldsStoreAnswers503Reported(t *testing.T) {
	q := service_tester.New(t)
	c := client_tester.GeneratedClient(t, base.Serve(t, q))
	held_store.New(t, q.Path)
	answer, e := c.PostDocumentWithResponse(
		t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: "notes",
			Path:       "alfa.md",
			Body:       "# Alfa\n\nzulu.\n",
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode())
	assert.FatalNotNil(t, answer.JSON503)
	assert.StringContains(t, "store busy", answer.JSON503.Error)
	assert.Count(t, 1, q.Reporter.Events())
}
