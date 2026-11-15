package chunk_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func Table(
	t *testing.T,
	allowance int,
) []chunk.Chunk {
	t.Helper()
	r := service_tester.NewReranker(t)
	r.SetAllowance(allowance)

	return chunk.Document(constant.FixtureTableDocument, "results.md", r)
}
