package preview_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/preview_chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func Window(
	t *testing.T,
	allowance int,
	body string,
) *preview_chunk.Chunk {
	t.Helper()
	r := service_tester.NewReranker(t)
	r.SetAllowance(allowance)

	return service_tester.NewService(
		t,
		service_tester.NewEmbedder(t),
		r,
	).Preview(
		"lines.md",
		body,
	).Chunks[0]
}
