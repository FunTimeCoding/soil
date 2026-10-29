package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/option"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/port_forward"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
)

func New(
	o *option.Server,
	l *logger.Logger,
	r face.Reporter,
) *Service {
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

	return &Service{
		database:     d,
		clusters:     clusters,
		guards:       guards,
		store:        store.New(d),
		reporter:     r,
		portForwards: map[string]*port_forward.Forward{},
	}
}
