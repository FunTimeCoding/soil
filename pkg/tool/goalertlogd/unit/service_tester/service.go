package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	l := logger.New(o.t.Context())
	d := relational.Open(l, "", o.Path)
	o.t.Cleanup(func() { relational.CloseMapper(d) })
	g := peer.New(
		constant.AlertmanagerPeer,
		o.Alertmanager.Address(),
		o.Alertmanager,
		l,
		o.Reporter,
	)
	o.t.Cleanup(g.Close)
	o.service = service.New(
		store.New(d),
		o.Alertmanager,
		g,
		o.retention,
		o.clock,
		o.notifier,
	)

	return o.service
}
