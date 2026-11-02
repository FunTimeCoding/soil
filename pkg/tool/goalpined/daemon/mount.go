package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	generated "github.com/funtimecoding/soil/pkg/tool/goalpined/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func (d *Daemon) Mount(
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(d.service, r),
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(d.service, r, t).Mount(g)
}
