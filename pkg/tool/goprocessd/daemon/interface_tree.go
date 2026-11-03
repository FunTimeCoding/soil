package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
	goprocessd "github.com/funtimecoding/soil/pkg/tool/goprocessd/server"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/service"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func interfaceTree(
	v *service.Service,
	r face.Reporter,
	t face.Recorder,
) http.Handler {
	return server.HandlerFromMux(
		server.NewStrictHandler(
			goprocessd.New(v, r),
			[]server.StrictMiddlewareFunc{
				web.RecordingMiddleware[server.StrictHandlerFunc](t),
			},
		),
		http.NewServeMux(),
	)
}
