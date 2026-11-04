package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/option"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward/registry"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
)

func New(
	o *option.Server,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	clusters := o.Clusters()
	guards := map[string]*peer.Guard{}

	for name, c := range clusters {
		guards[name] = peer.New(
			join.Colon(constant.ClusterPeer, name),
			c.Host(),
			c,
			l,
			r,
		)
	}

	forwards := registry.New()

	return &Daemon{
		database: d,
		clusters: clusters,
		guards:   guards,
		forwards: forwards,
		service:  service.New(clusters, guards, store.New(d), forwards, r),
		armed:    o.Arm,
	}
}
