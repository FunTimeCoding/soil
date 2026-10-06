package gosourced

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	s *service.Service,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	model_context.New(s, r, t).Mount(g)
}
