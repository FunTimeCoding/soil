package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/outpost"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/types/target"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/constant"
)

func (d *Daemon) addOutposts(o *option.Atlas) {
	var targets []*target.Target

	for _, host := range o.OutpostHosts {
		c := o.Outpost(host)
		g := peer.New(
			join.Colon(constant.Identity.Name(), host),
			c.Address(),
			c,
			d.logger,
			d.reporter,
		)
		d.peers = append(d.peers, c)
		d.guards = append(d.guards, g)
		targets = append(targets, target.New(host, c, g))
	}

	d.collectors = append(d.collectors, outpost.New(targets, d.logger))
}
