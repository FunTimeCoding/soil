package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
)

func (o *Tester) QueryVector(
	c *ollama.Client,
	query string,
) []float32 {
	o.t.Helper()
	result, e := embed.Single(o.t.Context(), c, query)
	assert.FatalOnError(o.t, e)

	return result
}
