package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/model_context"
	gosecurityd "github.com/funtimecoding/soil/pkg/tool/gosecurityd/server"
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
				gosecurityd.New(d.vulnerability, d.debian, d.flag, d.system, r),
				[]server.StrictMiddlewareFunc{
					web.RecordingMiddleware[server.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(d.vulnerability, d.debian, d.flag, d.system, r, t).Mount(
		g,
	)
	d.web.Mount(g)
}
