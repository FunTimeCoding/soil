package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	generated "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/server"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"time"
)

func (d *Daemon) Mount(
	r face.Reporter,
	hold time.Duration,
	t face.Recorder,
	g *guard.Mux,
) {
	g.TokenMount(
		webConstant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(
					d.roster,
					d.session,
					d.label,
					d.message,
					d.cruise,
					d.event,
					d.transcript,
					d.usage,
					d.search,
					d.logger,
					r,
					hold,
				),
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(
		d.roster,
		d.label,
		d.message,
		d.cruise,
		d.event,
		d.memory,
		d.usage,
		d.search,
		r,
		d.logger,
		t,
	).Mount(g)
	d.web.Mount(g)
}
