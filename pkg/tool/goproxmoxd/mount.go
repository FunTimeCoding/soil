package goproxmoxd

import (
	"github.com/funtimecoding/soil/pkg/face"
	proxFace "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	generated "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/server"
	proxmoxWeb "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/web"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mount(
	v proxFace.Service,
	b *proxmoxWeb.Server,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	g.TokenMount(
		constant.InterfacePath,
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
	model_context.New(v, r, t, version).Mount(g)
	b.Mount(g)
}
