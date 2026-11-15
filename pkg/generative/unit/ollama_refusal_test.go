package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/generative/unit/ollama_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestOllamaInputOverContextIsValidationError(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(
			`{"error":"the input length exceeds the context length"}`,
		),
	)
	_, e := ollama_tester.New(t, s.URL).Embed(t.Context(), []string{"alfa"})
	assert.True(t, validation.Is(e))
	assert.StringContains(
		t,
		"the input length exceeds the context length",
		e.Error(),
	)
}

func TestOllamaModelNotPulledIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(
			`{"error":"model \"nomic-embed-text\" not found, try pulling it first"}`,
		),
	)
	_, e := ollama_tester.New(t, s.URL).Embed(t.Context(), []string{"alfa"})
	assert.True(t, not_found.Is(e))
}

func TestFullOllamaQueueIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
		upstream_server.WithAnswer(
			`{"error":"server busy, please try again.  maximum pending requests exceeded"}`,
		),
	)
	_, e := ollama_tester.New(t, s.URL).Embed(t.Context(), []string{"alfa"})
	assert.True(t, unavailable.Is(e))
}

func TestOllamaServerErrorIsUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
		upstream_server.WithAnswer(
			`{"error":"llama runner process has terminated"}`,
		),
	)
	_, e := ollama_tester.New(t, s.URL).Embed(t.Context(), []string{"alfa"})
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "500", e.Error())
}
