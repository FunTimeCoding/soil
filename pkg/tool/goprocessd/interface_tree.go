package goprocessd

import (
	"github.com/funtimecoding/soil/pkg/face"
	generated "github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/server"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/service"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func interfaceTree(
	v *service.Service,
	r face.Reporter,
	t face.Recorder,
) http.Handler {
	return generated.HandlerFromMux(
		generated.NewStrictHandler(
			server.New(v, r),
			[]generated.StrictMiddlewareFunc{
				web.RecordingMiddleware[generated.StrictHandlerFunc](t),
			},
		),
		http.NewServeMux(),
	)
}
