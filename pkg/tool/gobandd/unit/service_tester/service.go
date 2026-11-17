package service_tester

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	l := logger.New(o.t.Context())
	p := pool.New(
		o.Inventory(),
		func(i *instance.Instance) (*peered.Peered, error) {
			c := o.instances[i.Name]

			return peered.New(
				c,
				peer.New(
					join.Colon(constant.InstancePeer, i.Name),
					i.Host,
					c,
					l,
					o.Reporter,
				),
			), nil
		},
	)
	o.t.Cleanup(p.Close)
	o.service = service.New(p, l, o.Reporter)

	if o.Netbox != nil {
		g := peer.New(
			gonetboxdConstant.Identity.Name(),
			o.Netbox.Address(),
			o.Netbox,
			l,
			o.Reporter,
		)
		o.t.Cleanup(g.Close)
		o.service = o.service.WithNetbox(o.Netbox, g)
	}

	return o.service
}
