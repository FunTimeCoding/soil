package goprometheusd

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/service"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	v *service.Service,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	model_context.New(v, r, t).Mount(g)
}
