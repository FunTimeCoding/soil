package service_tester

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/peered"
)

func (o *Tester) peers() *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered] {
	o.t.Helper()

	if o.pool != nil {
		return o.pool
	}

	l := logger.New(o.t.Context())
	o.pool = pool.New(
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
	o.t.Cleanup(o.pool.Close)

	return o.pool
}
