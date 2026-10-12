package gosecretd

import (
	"github.com/funtimecoding/soil/pkg/face"
	secret "github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
	generated "github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/service"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	c secret.SecretSource,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	s := service.New(c)
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
	model_context.New(s, r, t).Mount(g)
}
