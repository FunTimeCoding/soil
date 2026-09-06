package gomemoryd

import (
	"github.com/funtimecoding/soil/pkg/face"
	generated "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	memoryWeb "github.com/funtimecoding/soil/pkg/tool/gomemoryd/web"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	v *service.Service,
	u *memoryWeb.Server,
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
