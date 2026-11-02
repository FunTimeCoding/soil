package embedder

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"testing"
)

func assertEmbedded(
	t *testing.T,
	e face.Embedder,
	name string,
	body string,
) {
	t.Helper()

	for _, c := range chunk.Document(body, name, newReranker(t)) {
		_, f := embed.Single(t.Context(), e, c.Text)
		assert.FatalOnError(t, f)
	}
}
