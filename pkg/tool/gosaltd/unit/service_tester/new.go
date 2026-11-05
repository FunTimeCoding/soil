package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/migrate"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/service"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func New(
	t *testing.T,
	o *option.Salt,
) (*service.Service, *runner.Runner) {
	t.Helper()
	m := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m, constant.RunsTable)
	s := store.New(m, constant.RunsTable)
	c := o.Master()
	t.Cleanup(c.Close)
	r := memory.New()
	p := reaper.New(r, prometheus.NewRegistry())
	t.Cleanup(p.Close)
	l := logger.New(t.Context())
	g := peer.New(constant.SaltPeer, c.Address(), c, l, r)
	t.Cleanup(g.Close)
	u := runner.New(o, s, l, r, p, c, g)
	t.Cleanup(
		func() {
			u.Close()
			<-u.Finished()
		},
	)

	return service.New(s, u, c, g), u
}
