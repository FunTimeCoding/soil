package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func (d *Daemon) MountLocal(
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	g.Open(
		route.Get(webConstant.ReachablePath),
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			w.WriteHeader(http.StatusNoContent)
		},
	)
	g.OpenMount(webConstant.InterfacePath, interfaceTree(d.service, r, t))
}
