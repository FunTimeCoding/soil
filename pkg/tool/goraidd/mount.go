package goraidd

import (
	"github.com/funtimecoding/soil/pkg/face"
	generated "github.com/funtimecoding/soil/pkg/tool/goraidd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/server"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	raidWeb "github.com/funtimecoding/soil/pkg/tool/goraidd/web"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	s *store.Store,
	outputPath string,
	u *raidWeb.Server,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(s, outputPath, r),
				[]generated.StrictMiddlewareFunc{
					web.RecordingMiddleware[generated.StrictHandlerFunc](t),
				},
			),
			http.NewServeMux(),
		),
	)
	u.Mount(g)
}
