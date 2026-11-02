package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/model_context"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func (d *Daemon) Mount(
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	model_context.New(d.service, d.inventory, d.selection, r, t).Mount(g)
}
