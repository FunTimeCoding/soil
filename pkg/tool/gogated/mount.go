package gogated

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/server"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	v *service.Service,
	administration *web.Server,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	server.New(v, r).Mount(g)
	administration.Mount(g)
	model_context.New(v, r, t).Mount(g)
}
