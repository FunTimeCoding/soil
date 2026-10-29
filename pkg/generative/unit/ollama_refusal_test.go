package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"net/http"
	"testing"
)

func TestAnOllamaInputOverTheContextIsAValidationError(t *testing.T) {
	e := ollamaEmbedding(
		t,
		http.StatusBadRequest,
		`{"error":"the input length exceeds the context length"}`,
	)
	assert.True(t, validation.Is(e))
	assert.StringContains(
		t,
		"the input length exceeds the context length",
		e.Error(),
	)
}

func TestAnOllamaModelNotPulledIsNotFound(t *testing.T) {
	e := ollamaEmbedding(
		t,
		http.StatusNotFound,
		`{"error":"model \"nomic-embed-text\" not found, try pulling it first"}`,
	)
	assert.True(t, not_found.Is(e))
}

func TestAFullOllamaQueueIsUnavailable(t *testing.T) {
	e := ollamaEmbedding(
		t,
		http.StatusServiceUnavailable,
		`{"error":"server busy, please try again.  maximum pending requests exceeded"}`,
	)
	assert.True(t, unavailable.Is(e))
}

func TestAnOllamaServerErrorIsUnexpected(t *testing.T) {
	e := ollamaEmbedding(
		t,
		http.StatusInternalServerError,
		`{"error":"llama runner process has terminated"}`,
	)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "500", e.Error())
}
