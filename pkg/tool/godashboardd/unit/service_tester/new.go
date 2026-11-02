package service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/option"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/store"
	"testing"
)

func New(
	t *testing.T,
	o *option.Dashboard,
	r face.Reporter,
) *service.Service {
	t.Helper()
	l := logger.New(t.Context())
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	t.Cleanup(func() { relational.CloseMapper(d) })
	p := o.Prometheus()
	g := peer.New(constant.PrometheusPeer, p.Address(), p, l, r)
	t.Cleanup(g.Close)
	a := o.Argocd()
	var argocdGuard *peer.Guard

	if a != nil {
		argocdGuard = peer.New(
			constant.ArgocdPeer,
			o.Board.Connection.Argocd.Host,
			a,
			l,
			r,
		)
		t.Cleanup(argocdGuard.Close)
	}

	u := o.Usage()
	var usageGuard *peer.Guard

	if u != nil {
		usageGuard = peer.New(
			constant.NextcloudPeer,
			o.Board.Connection.Nextcloud.Host,
			u,
			l,
			r,
		)
		t.Cleanup(usageGuard.Close)
	}

	return service.New(
		o.Board,
		store.New(d),
		p,
		g,
		a,
		argocdGuard,
		u,
		usageGuard,
		notifier.New(),
	)
}
