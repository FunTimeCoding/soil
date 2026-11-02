package service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/service"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store"
	"testing"
)

func New(t *testing.T) *service.Service {
	t.Helper()
	m := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m)
	n := notifier.New()

	return service.New(store.New(m, n), n)
}
