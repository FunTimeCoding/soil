package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"testing"
)

func NewStore(t *testing.T) *store.Store {
	t.Helper()
	m := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m)

	return store.New(m)
}
