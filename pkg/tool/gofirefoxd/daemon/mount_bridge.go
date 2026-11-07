package daemon

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func (d *Daemon) MountBridge(g *guard.Mux) {
	g.OpenMount(webConstant.RootPath, d.client)
}
