package daemon

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/option"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web"
	gitlabWorker "github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	o *option.Gitlab,
	l *logger.Logger,
	r face.Reporter,
	y *prometheus.Registry,
) *Daemon {
	p := pool.New(o.Inventory, peeredBuilder(o, l, r))

	for _, i := range o.Inventory.Instances {
		_, e := p.Client(i.Name)
		errors.PanicOnError(e)
	}

	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	s := service.New(p, o.Inventory)
	result := &Daemon{
		inventory: o.Inventory,
		pool:      p,
		selection: selection.New(),
		service:   s,
		web:       web.New(s, o.Inventory, r),
		worker: gitlabWorker.New(
			s.Refresh,
			s.Entries,
			constant.PollInterval,
			y,
			l,
			r,
		),
		database: d,
	}

	if c := o.Connector(); c != nil {
		result.connector = c
		result.rosterGuard = peer.New(constant.RosterPeer, c.Address(), c, l, r)
		s.WithWatches(store.New(d), c, result.rosterGuard, l, r)
		result.clocks = append(
			result.clocks,
			worker.New(s.PollWatches, l, r, o.WatchInterval),
			worker.New(s.SweepRoster, l, r, o.SweepInterval),
		)
	}

	return result
}
