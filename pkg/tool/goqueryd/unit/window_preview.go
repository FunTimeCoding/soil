package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/preview_chunk"
	"testing"
)

func windowPreview(
	t *testing.T,
	allowance int,
	body string,
) *preview_chunk.Chunk {
	t.Helper()
	r := newReranker(t)
	r.SetAllowance(allowance)

	return newService(t, newEmbedder(t), r).Preview("lines.md", body).Chunks[0]
}
