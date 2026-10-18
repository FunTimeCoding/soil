package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"testing"
)

func queryVector(
	t *testing.T,
	o *ollama.Client,
	query string,
) []float32 {
	t.Helper()
	result, e := embed.Single(t.Context(), o, query)
	assert.FatalOnError(t, e)

	return result
}
