package goalertlogd

import (
	"github.com/funtimecoding/soil/pkg/face"
	generated "github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/server"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	alertWeb "github.com/funtimecoding/soil/pkg/tool/goalertlogd/web"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	s *store.Store,
	w *worker.Worker,
	u *alertWeb.Server,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(s, w, r),
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(s, w, r, t, version).Mount(g)
	u.Mount(g)
}
