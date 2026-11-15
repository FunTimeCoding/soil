package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/held_store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"net/http"
	"testing"
)

func TestPushWhileAnotherWriterHoldsStoreAnswers503Reported(t *testing.T) {
	r := memory.New()
	o := service_tester.NewOption(
		t,
		service_tester.NewEmbedder(t),
		service_tester.NewReranker(t),
	)
	c := client_tester.Rest(t, base.Serve(t, o, r))
	held_store.New(t, o.LitePath)
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
	assert.Count(t, 1, r.Events())
}
