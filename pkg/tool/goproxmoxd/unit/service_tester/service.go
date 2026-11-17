package service_tester

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	i := o.Inventory()
	l := logger.New(o.t.Context())
	p := pool.New(
		i,
		func(n *instance.Instance) (*peered.Peered, error) {
			c := o.Instance(n.Name).Hypervisor

			return peered.New(
				c,
				peer.New(
					join.Colon(constant.InstancePeer, n.Name),
					n.Host,
					c,
					l,
					o.Reporter,
				),
			), nil
		},
	)
	o.t.Cleanup(p.Close)
	s := pool.New(i, o.shell)
	o.t.Cleanup(s.Close)
	o.service = service.New(p, s, i, l, o.Reporter)

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
