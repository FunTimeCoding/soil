package daemon

import "github.com/funtimecoding/soil/pkg/web/guard"

func (d *Daemon) Mount(g *guard.Mux) {
	d.web.Mount(g)
}
