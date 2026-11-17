package preview_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/preview_chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func Window(
	t *testing.T,
	tokenLimit int,
	body string,
) *preview_chunk.Chunk {
	t.Helper()
	q := service_tester.New(t)
	q.Reranker.SetTokenLimit(tokenLimit)

	return q.Service().Preview("lines.md", body).Chunks[0]
}
