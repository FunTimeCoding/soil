package service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/service"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service == nil {
		m := lite.NewMemory()
		o.t.Cleanup(func() { relational.CloseMapper(m) })
		migrate.AutoMigrate(m)
		n := notifier.New()
		o.service = service.New(store.New(m, n), n)
	}

	return o.service
}
