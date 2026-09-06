package goraidparsed

import (
	"github.com/funtimecoding/soil/pkg/face"
	generated "github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	v *server.Server,
	t face.Recorder,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				v,
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
}
