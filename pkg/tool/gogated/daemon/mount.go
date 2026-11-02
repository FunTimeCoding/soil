package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/server"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func (d *Daemon) Mount(
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	server.New(d.service, r).Mount(g)
	d.web.Mount(g)
	model_context.New(d.service, r, t).Mount(g)
}
