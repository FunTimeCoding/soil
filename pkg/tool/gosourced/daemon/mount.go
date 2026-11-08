package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context"
	gosourced "github.com/funtimecoding/soil/pkg/tool/gosourced/server"
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
				gosourced.New(d.census, d.diff, d.module, r),
				[]server.StrictMiddlewareFunc{
					web.RecordingMiddleware[server.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(
		d.file,
		d.module,
		d.packages,
		d.pattern,
		d.reference,
		d.symbol,
		r,
		t,
	).Mount(g)
}
