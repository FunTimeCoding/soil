package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/process"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	goprocessdConstant "github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
)

func (d *Daemon) addProcess(o *option.Atlas) {
	c := o.Process()
	g := peer.New(
		goprocessdConstant.Identity.Name(),
		c.Address(),
		c,
		d.logger,
		d.reporter,
	)
	d.peers = append(d.peers, c)
	d.guards = append(d.guards, g)
	d.collectors = append(d.collectors, process.New(c, g, o.ProcessPlace))
}
