package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"net/http"
	"testing"
)

func TestSearchWhileEmbedderIsDownAnswers503WithoutReport(t *testing.T) {
	m := service_tester.NewEmbedder(t)
	m.Fail(outage_tester.RefusedDial())
	r := memory.New()
	c := client_tester.Rest(
		t,
		base.Serve(
			t,
			service_tester.NewOption(t, m, service_tester.NewReranker(t)),
			r,
		),
	)
	answer, e := c.GetSearchWithResponse(
		t.Context(),
		&client.GetSearchParams{Query: "zulu"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode())
	assert.FatalNotNil(t, answer.JSON503)
	assert.StringContains(t, "embedder unreachable since", answer.JSON503.Error)
	assert.Count(t, 0, r.Events())
}

func TestPushWhileEmbedderIsDownAnswersItsWarning(t *testing.T) {
	m := service_tester.NewEmbedder(t)
	m.Fail(outage_tester.RefusedDial())
	c := client_tester.Rest(
		t,
		base.Serve(
			t,
			service_tester.NewOption(t, m, service_tester.NewReranker(t)),
			memory.New(),
		),
	)
	answer, e := c.PostDocumentWithResponse(
		t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: "notes",
			Path:       "alfa.md",
			Body:       "# Alfa\n\nzulu.\n",
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, answer.StatusCode())
	assert.FatalNotNil(t, answer.JSON200)
	assert.FatalNotNil(t, answer.JSON200.Warning)
	assert.StringContains(
		t,
		"stored; embedding pending",
		*answer.JSON200.Warning,
	)
}
