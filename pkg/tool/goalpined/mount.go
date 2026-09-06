package goalpined

import (
	"github.com/funtimecoding/soil/pkg/alpine/package_server"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/goalpined/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	p *package_server.Server,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	inner := http.NewServeMux()
	p.Mount(inner)
	generated.HandlerFromMux(
		generated.NewStrictHandler(
			server.New(r),
			[]generated.StrictMiddlewareFunc{
				web.RecordingMiddleware[generated.StrictHandlerFunc](t),
			},
		),
		inner,
	)
	g.TokenMount(constant.Slash, inner)
	model_context.New(r, t, version).Mount(g)
}
