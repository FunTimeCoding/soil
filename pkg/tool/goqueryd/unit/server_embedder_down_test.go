package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"net/http"
	"testing"
)

func TestASearchWhileTheEmbedderIsDownAnswers503WithoutAReport(t *testing.T) {
	m := newEmbedder(t)
	m.Fail(refusedConnection())
	v := newService(t, m, newReranker(t))
	r := memory.New()
	c := mountedClient(t, v, r)
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

func TestAPushWhileTheEmbedderIsDownAnswersItsWarning(t *testing.T) {
	m := newEmbedder(t)
	m.Fail(refusedConnection())
	v := newService(t, m, newReranker(t))
	c := mountedClient(t, v, memory.New())
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
