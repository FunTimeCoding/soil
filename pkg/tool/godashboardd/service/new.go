package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/option"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/store"
)

func New(
	o *option.Dashboard,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	p := o.Prometheus()
	result := &Service{
		board:           o.Board,
		database:        d,
		store:           store.New(d),
		prometheus:      p,
		prometheusGuard: peer.New(constant.PrometheusPeer, p.Address(), p, l, r),
		argocd:          o.Argocd(),
		usage:           o.Usage(),
		authorization:   authorizationClient(o).WithReporter(r),
		notifier:        notifier.New(),
		labels:          map[string]bool{},
		values:          map[string][]string{},
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

	for _, entry := range o.Board.Entries() {
		result.labels[entry.Label] = true
	}

	return result
}
