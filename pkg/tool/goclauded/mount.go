package goclauded

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	generated "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web"
	soilWeb "github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	v *service.Service,
	u *web.Server,
	l *logger.Logger,
	r face.Reporter,
	harborPath string,
	sessionExportPath string,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(v, l, r, harborPath, sessionExportPath),
				[]generated.StrictMiddlewareFunc{
					soilWeb.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(v, r, l, t, version).Mount(g)
	u.Mount(g)
}
