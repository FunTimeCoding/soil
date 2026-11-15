package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"testing"
)

func QueryVector(
	t *testing.T,
	o *ollama.Client,
	query string,
) []float32 {
	t.Helper()
	result, e := embed.Single(t.Context(), o, query)
	assert.FatalOnError(t, e)

	return result
}
