package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"testing"
)

func tableChunks(
	t *testing.T,
	allowance int,
) []chunk.Chunk {
	t.Helper()
	r := newReranker(t)
	r.SetAllowance(allowance)

	return chunk.Document(constant.FixtureTableDocument, "results.md", r)
}
