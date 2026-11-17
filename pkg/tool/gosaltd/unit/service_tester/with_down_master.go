package service_tester

import (
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/salt/mock_client"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
)

func WithDownMaster() Option {
	return func(o *Tester) {
		c := mock_client.New(o.t)
		c.Fail(c.Reachable, outage_tester.RefusedDial())
		c.Fail(c.Keys, outage_tester.RefusedDial())
		c.Fail(c.Minions, outage_tester.RefusedDial())
		o.master = func() face.SaltClient { return c }
	}
}
