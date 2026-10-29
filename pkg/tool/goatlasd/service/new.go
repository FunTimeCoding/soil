package service

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
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func New(
	o *option.Atlas,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	n := o.Netbox()
	ng := peer.New(netbox.Identity.Name(), n.Address(), n, l, r)
	k := o.Kubernetes()
	kg := peer.New(constant.SourceKubernetes, k.Configuration().Host, k, l, r)
	result := &Service{
		database:      d,
		store:         store.New(d),
		authorization: authorizationClient(o).WithReporter(r),
		netbox:        n,
		netboxGuard:   ng,
		guards:        []*peer.Guard{kg},
		collectors: []atlas.Collector{
			kubernetes.New(k, kg),
			label.New(n, ng),
		},
		retention: o.Retention,
	}

	if o.Process != nil {
		result.addProcess(o, l, r)
	}

	if o.Outpost != nil && len(o.OutpostHosts) > 0 {
		result.addOutposts(o, l, r)
	}

	if o.Lease != nil {
		result.addLease(o, l, r)
	}

	return result
}
