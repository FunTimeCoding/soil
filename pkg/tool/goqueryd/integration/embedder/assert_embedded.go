package embedder

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
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

	for _, c := range chunk.Document(body, name, mock_reranker.New()) {
		_, f := embed.Single(t.Context(), e, c.Text)
		assert.FatalOnError(t, f)
	}
}
