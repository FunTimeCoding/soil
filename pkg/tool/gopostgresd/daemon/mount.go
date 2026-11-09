package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context"
	gopostgresd "github.com/funtimecoding/soil/pkg/tool/gopostgresd/server"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func (d *Daemon) Mount(
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	g.TokenMount(
		webConstant.InterfacePath,
		server.HandlerFromMux(
			server.NewStrictHandler(
				gopostgresd.New(d.catalog, d.query, d.backup, d.inventory, r),
				[]server.StrictMiddlewareFunc{
					web.RecordingMiddleware[server.StrictHandlerFunc](t),
					gopostgresd.ExtendedWrite,
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(
		d.catalog,
		d.query,
		d.backup,
		d.inventory,
		d.selection,
		r,
		t,
	).Mount(g)
}
