package gosentryd

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/service"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	v *service.Service,
	host string,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	model_context.New(v, host, r, t).Mount(g)
}
