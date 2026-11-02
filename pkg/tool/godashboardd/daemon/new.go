package daemon

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
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/web"
)

func New(
	o *option.Dashboard,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	p := o.Prometheus()
	result := &Daemon{
		database:        d,
		prometheus:      p,
		prometheusGuard: peer.New(constant.PrometheusPeer, p.Address(), p, l, r),
		argocd:          o.Argocd(),
		usage:           o.Usage(),
		authorization:   authorizationClient(o).WithReporter(r),
		notifier:        notifier.New(),
		logger:          l,
		reporter:        r,
	}

	if result.argocd != nil {
		result.argocdGuard = peer.New(
			constant.ArgocdPeer,
			o.Board.Connection.Argocd.Host,
			result.argocd,
			l,
			r,
		)
	}

	if result.usage != nil {
		result.usageGuard = peer.New(
			constant.NextcloudPeer,
			o.Board.Connection.Nextcloud.Host,
			result.usage,
			l,
			r,
		)
	}

	result.service = service.New(
		o.Board,
		store.New(d),
		p,
		result.prometheusGuard,
		result.argocd,
		result.argocdGuard,
		result.usage,
		result.usageGuard,
		result.notifier,
	)
	result.web = web.New(
		result.service,
		o.Board,
		result.authorization,
		result.notifier,
		r,
	)

	return result
}
