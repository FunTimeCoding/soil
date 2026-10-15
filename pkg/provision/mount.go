package provision

import (
	"github.com/funtimecoding/soil/pkg/face"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
	generated "github.com/funtimecoding/soil/pkg/provision/generated/server"
	"github.com/funtimecoding/soil/pkg/provision/server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	n provisionFace.Runner,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(n, r),
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
}
