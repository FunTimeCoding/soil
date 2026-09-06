package gosublimed

import (
	"github.com/funtimecoding/soil/pkg/face"
	sublime "github.com/funtimecoding/soil/pkg/tool/gosublimed/face"
	generated "github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	c sublime.SublimeSource,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(c),
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(c, r, t, version).Mount(g)
}
