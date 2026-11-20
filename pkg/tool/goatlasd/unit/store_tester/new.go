package store_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	m := lite.NewMemory()
	migrate.AutoMigrate(m)

	return &Tester{Store: store.New(m), t: t}
}
