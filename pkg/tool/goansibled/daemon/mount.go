package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/provision"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/model_context"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func (d *Daemon) Mount(
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	provision.Mount(d.service, r, t, g)
	model_context.New(d.service, r, t).Mount(g)
}
