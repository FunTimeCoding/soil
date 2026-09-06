package gosproutd

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	v *service.Service,
	u *web.Server,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	model_context.New(v, r, t, version).Mount(g)
	u.Mount(g)
}
