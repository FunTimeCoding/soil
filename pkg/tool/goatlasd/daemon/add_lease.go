package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
)

func (d *Daemon) addLease(o *option.Atlas) {
	c := o.Lease()
	g := peer.New(
		constant.Identity.Name(),
		c.Address(),
		c,
		d.logger,
		d.reporter,
	)
	d.peers = append(d.peers, c)
	d.guards = append(d.guards, g)
	d.lease = lease.New(c, g)
}
