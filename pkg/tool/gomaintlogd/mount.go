package gomaintlogd

import (
	"github.com/funtimecoding/soil/pkg/face"
	generated "github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
	maintenanceWeb "github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	s *store.Store,
	v *maintenanceWeb.Server,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(s, r),
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(s, r, t, version).Mount(g)
	v.Mount(g)
}
