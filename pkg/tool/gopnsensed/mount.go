package gopnsensed

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/opnsense"
	generated "github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	c *opnsense.Client,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(c, r),
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(c, r, t, version).Mount(g)
}
