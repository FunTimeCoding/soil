package goqueryd

import (
	"github.com/funtimecoding/soil/pkg/face"
	generated "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	queryWeb "github.com/funtimecoding/soil/pkg/tool/goqueryd/web"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	v *service.Service,
	u *queryWeb.Server,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(v, r),
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(v, r, t, version).Mount(g)
	u.Mount(g)
}
