package gotelemetryd

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/web"
	soilWeb "github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	v *service.Service,
	u *web.Server,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	h := generated.HandlerFromMux(
		generated.NewStrictHandler(
			server.New(v, r),
			[]generated.StrictMiddlewareFunc{
				soilWeb.RecordingMiddleware[generated.StrictHandlerFunc](
					t,
					constant.IngestOperation,
				),
			},
		),
		http.NewServeMux(),
	)
	g.TokenMount(webConstant.InterfacePath, h)
	g.OpenMount(constant.IngestPattern, h)
	model_context.New(v, r, t).Mount(g)
	u.Mount(g)
}
