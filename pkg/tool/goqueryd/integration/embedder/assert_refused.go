package embedder

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"testing"
)

func assertRefused(
	t *testing.T,
	e face.Embedder,
	name string,
	body string,
) {
	t.Helper()
	chunks := chunk.Document(body, name, mock_reranker.New())
	assert.Count(t, 1, chunks)
	_, f := embed.Single(t.Context(), e, chunks[0].Text)
	assert.True(t, validation.Is(f))
}
