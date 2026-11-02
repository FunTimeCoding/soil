package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/server"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func (d *Daemon) Mount(
	r face.Reporter,
	g *guard.Mux,
) {
	h := generated.HandlerFromMux(
		generated.NewStrictHandler(
			server.New(d.service, r),
			[]generated.StrictMiddlewareFunc{
				web.RecordingMiddleware[generated.StrictHandlerFunc](
					d.service,
					constant.IngestOperation,
				),
			},
		),
		http.NewServeMux(),
	)
	g.TokenMount(webConstant.InterfacePath, h)
	g.OpenMount(constant.IngestPattern, h)
	model_context.New(d.service, r, d.service).Mount(g)
	d.web.Mount(g)
}
