package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"testing"
)

func NewStore(t *testing.T) *store.Store {
	t.Helper()
	m := lite.NewMemory()
	migrate.AutoMigrate(m)

	return store.New(m)
}
