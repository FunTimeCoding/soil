package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"testing"
	"time"
)

func NewStore(
	t *testing.T,
	clock func() time.Time,
) *store.Store {
	t.Helper()
	d := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(d) })
	migrate.AutoMigrate(d)

	return store.New(d, clock)
}
