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
	v *server.Server,
	administration *web.Server,
	s *service.Service,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	v.Mount(g)
	administration.Mount(g)
	model_context.New(s, r, t).Mount(g)
}
