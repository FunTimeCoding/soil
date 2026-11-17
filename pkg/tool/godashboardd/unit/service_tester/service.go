package service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/store"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	l := logger.New(o.t.Context())
	d := relational.Open(l, "", o.litePath)
	migrate.AutoMigrate(d)
	o.t.Cleanup(func() { relational.CloseMapper(d) })
	g := peer.New(
		constant.PrometheusPeer,
		o.Prometheus.Address(),
		o.Prometheus,
		l,
		o.Reporter,
	)
	o.t.Cleanup(g.Close)
	var argocdGuard *peer.Guard

	if o.Argocd != nil {
		argocdGuard = peer.New(
			constant.ArgocdPeer,
			o.board.Connection.Argocd.Host,
			o.Argocd,
			l,
			o.Reporter,
		)
		o.t.Cleanup(argocdGuard.Close)
	}

	var usageGuard *peer.Guard

	if o.Usage != nil {
		usageGuard = peer.New(
			constant.NextcloudPeer,
			o.board.Connection.Nextcloud.Host,
			o.Usage,
			l,
			o.Reporter,
		)
		o.t.Cleanup(usageGuard.Close)
	}

	o.service = service.New(
		o.board,
		store.New(d),
		o.Prometheus,
		g,
		o.argocd(),
		argocdGuard,
		o.usage(),
		usageGuard,
		notifier.New(),
	)

	return o.service
}
