package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func rechunkService(
	t *testing.T,
	position int,
) (*service.Service, *store.Store) {
	t.Helper()
	o := service_tester.NewOption(t, newEmbedder(t), newReranker(t))
	s := store_tester.OpenFileStore(t, o.LitePath)
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

	return service_tester.New(t, o, memory.New()), s
}
