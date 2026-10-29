package gomonitord

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	monitor "github.com/funtimecoding/soil/pkg/tool/gomonitord/face"
	generated "github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/server"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/stream"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func Mount(
	v monitor.Monitor,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	g.TokenMount(
		webConstant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(v, r),
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	g.Token(route.Get(constant.StreamPath), stream.Claims(v, r))
}
