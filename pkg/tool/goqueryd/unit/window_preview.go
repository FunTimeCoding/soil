package unit

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service/preview_chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func windowPreview(
	t *testing.T,
	allowance int,
	body string,
) *preview_chunk.Chunk {
	t.Helper()
	r := mock_reranker.New()
	r.SetAllowance(allowance)

	return service.New(
		store_tester.OpenTestStore(t),
		ollama.New(),
		r,
	).Preview("lines.md", body).Chunks[0]
}
