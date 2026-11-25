package service_tester

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	l := logger.New(o.t.Context())
	i := o.Inventory()
	p := pool.New(
		i,
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
	o.service = service.New(p, i)

	if c := o.connector; c != nil {
		d := lite.New(l, o.Path)
		o.t.Cleanup(func() { relational.CloseMapper(d) })
		migrate.AutoMigrate(d)
		g := peer.New(constant.RosterPeer, c.Address(), c, l, o.Reporter)
		o.t.Cleanup(g.Close)
		o.service.WithWatches(store.New(d), c, g, l, o.Reporter)
	}

	return o.service
}
