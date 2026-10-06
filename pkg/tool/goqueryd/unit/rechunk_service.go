package unit

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func rechunkService(
	t *testing.T,
	position int,
) (*service.Service, *store.Store) {
	t.Helper()
	s := store_tester.OpenTestStore(t)
	t.Cleanup(s.Close)
	d := t.TempDir()
	store_tester.WriteFixture(t, d, "alfa.md", "one two three\n")
	s.AddCollection("test", d, constant.DefaultGlob)
	s.Index("test")
	s.InsertEmbedding(
		store.HashContent("one two three\n"),
		0,
		position,
		[]float32{1},
		"2024-01-01T00:00:00Z",
	)

	return service.New(s, ollama.New(), mock_reranker.New()), s
}
