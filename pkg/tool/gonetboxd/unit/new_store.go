package unit

import (
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
	"testing"
)

func newStore(t *testing.T) *store.Store {
	t.Helper()
	d := lite.NewMemory()
	migrate.AutoMigrate(d)
	t.Cleanup(func() { relational.CloseMapper(d) })

	return store.New(d)
}
