package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
)

func tableChunks(allowance int) []chunk.Chunk {
	r := mock_reranker.New()
	r.SetAllowance(allowance)

	return chunk.Document(constant.FixtureTableDocument, "results.md", r)
}
