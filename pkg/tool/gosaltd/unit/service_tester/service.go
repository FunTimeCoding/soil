package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/migrate"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/service"
	"github.com/prometheus/client_golang/prometheus"
)

func (o *Tester) Service() (*service.Service, *runner.Runner) {
	o.t.Helper()

	if o.service != nil {
		return o.service, o.runner
	}

	m := lite.NewMemory()
	o.t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m, constant.RunsTable)
	s := store.New(m, constant.RunsTable)
	c := o.master()
	o.t.Cleanup(c.Close)
	p := reaper.New(o.Reporter, prometheus.NewRegistry())
	o.t.Cleanup(p.Close)
	l := logger.New(o.t.Context())
	g := peer.New(constant.SaltPeer, c.Address(), c, l, o.Reporter)
	o.t.Cleanup(g.Close)
	o.runner = runner.New(o.daemonOption(), s, l, o.Reporter, p, c, g)
	o.t.Cleanup(
		func() {
			o.runner.Close()
			<-o.runner.Finished()
		},
	)
	o.service = service.New(s, o.runner, c, g)

	return o.service, o.runner
}
