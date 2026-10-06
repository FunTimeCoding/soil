package gojellyfind

import (
	"github.com/funtimecoding/soil/pkg/face"
	jellyfin "github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
	generated "github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	c jellyfin.JellyfinSource,
	r face.Reporter,
	t face.Recorder,
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
	model_context.New(c, r, t).Mount(g)
}
