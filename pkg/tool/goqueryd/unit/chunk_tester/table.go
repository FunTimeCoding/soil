package chunk_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func Table(
	t *testing.T,
	tokenLimit int,
) []chunk.Chunk {
	t.Helper()
	r := service_tester.New(t).Reranker
	r.SetTokenLimit(tokenLimit)

	return chunk.Document(constant.FixtureTableDocument, "results.md", r)
}
