package goatlasd

import (
	"github.com/funtimecoding/soil/pkg/face"
	generated "github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	atlasWeb "github.com/funtimecoding/soil/pkg/tool/goatlasd/web"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	s *store.Store,
	u *atlasWeb.Server,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	u.Mount(g)
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
}
