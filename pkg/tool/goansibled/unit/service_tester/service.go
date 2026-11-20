package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/migrate"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/runner"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/service"
	"github.com/prometheus/client_golang/prometheus"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	p := o.Option
	l := logger.New(o.t.Context())
	m := relational.Open(l, p.PostgresLocator, p.LitePath)
	o.t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m, constant.RunsTable)
	s := store.New(m, constant.RunsTable)
	k := reaper.New(o.Reporter, prometheus.NewRegistry())
	k.Start()
	o.t.Cleanup(k.Close)
	r := runner.New(p, s, l, o.Reporter, k)

	if p.Inventory != nil {
		c := p.Inventory()
		o.t.Cleanup(c.Close)
		r.WithInventory(c)
	}

	r.Start()
	o.t.Cleanup(
		func() {
			r.Close()
			<-r.Finished()
		},
	)
	o.service = service.New(s, r)

	return o.service
}
