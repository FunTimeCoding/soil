package rechunk_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func New(
	t *testing.T,
	position int,
) (*service.Service, *store.Store) {
	t.Helper()
	q := service_tester.New(t)
	s := store_tester.OpenFileStore(t, q.Path)
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

	return q.Service(), s
}
