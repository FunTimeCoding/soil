package golinkaced

import (
	"github.com/funtimecoding/soil/pkg/face"
	linkace "github.com/funtimecoding/soil/pkg/tool/golinkaced/face"
	generated "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/server"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/service"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	c linkace.LinkAceSource,
	s *service.Service,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(c, s, r),
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	model_context.New(c, s, r, t).Mount(g)
}
