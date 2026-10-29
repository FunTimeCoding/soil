package unit

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store"
	"testing"
)

func newStore(
	t *testing.T,
	n *notifier.Notifier,
) *store.Store {
	t.Helper()
	m := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m)

	return store.New(m, n)
}
