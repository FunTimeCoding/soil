package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	p := o.daemonOption()
	l := logger.New(o.t.Context())
	m := lite.New(l, o.Path)
	o.t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m)
	s := store.New(m)
	c := o.client
	g := peer.New(constant.MattermostPeer, c.Host(), c, l, o.Reporter)
	o.t.Cleanup(g.Close)
	o.service = service.New(
		s,
		c,
		g,
		l,
		o.Reporter,
		p.DebounceWindow,
		p.PurgeWindow,
	)

	if k := o.connector; k != nil {
		h := peer.New(constant.RosterPeer, k.Address(), k, l, o.Reporter)
		o.t.Cleanup(h.Close)
		w := watcher.New(c, g, s, l, o.Reporter).WithBackoff(
			p.ReconnectMinimum,
			p.ReconnectMaximum,
		)
		o.t.Cleanup(w.Close)
		o.service.WithSubscription(
			notifier.New(k, constant.Identity.Name(), h),
			w,
		)
	}

	if p.Monitor.Enabled {
		o.service.WithMonitor(monitor.New(c, g, p.Monitor, l))
	}

	return o.service
}
