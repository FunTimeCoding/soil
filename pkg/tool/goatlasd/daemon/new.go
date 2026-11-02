package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/kubernetes"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/label"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	atlas "github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/worker"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	o *option.Atlas,
	l *logger.Logger,
	r face.Reporter,
	y *prometheus.Registry,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	n := o.Netbox()
	ng := peer.New(netbox.Identity.Name(), n.Address(), n, l, r)
	k := o.Kubernetes()
	kg := peer.New(constant.SourceKubernetes, k.Configuration().Host, k, l, r)
	result := &Daemon{
		database:      d,
		authorization: authorizationClient(o).WithReporter(r),
		netbox:        n,
		netboxGuard:   ng,
		guards:        []*peer.Guard{kg},
		collectors: []atlas.Collector{
			kubernetes.New(k, kg),
			label.New(n, ng),
		},
		logger:   l,
		reporter: r,
	}

	if o.Process != nil {
		result.addProcess(o)
	}

	if o.Outpost != nil && len(o.OutpostHosts) > 0 {
		result.addOutposts(o)
	}

	if o.Lease != nil {
		result.addLease(o)
	}

	result.service = service.New(
		store.New(d),
		n,
		ng,
		result.collectors,
		result.lease,
		o.Retention,
	)
	result.web = web.New(result.service, result.authorization)
	result.clock = worker.New(result.service.Collect, o.Interval, l, y, r)

	return result
}
