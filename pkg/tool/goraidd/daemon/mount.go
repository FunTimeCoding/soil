package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/server"
	goraidd "github.com/funtimecoding/soil/pkg/tool/goraidd/server"
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
				goraidd.New(d.service, r),
				[]server.StrictMiddlewareFunc{
					web.RecordingMiddleware[server.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	d.web.Mount(g)
}
