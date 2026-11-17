package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/mock_system"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	m := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m)

	return &Tester{
		Store:    store.New(m),
		System:   mock_system.New(t),
		Reporter: memory.New(),
		mapper:   m,
		t:        t,
	}
}
