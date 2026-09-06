package gokubernetesd

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	s *service.Service,
	readOnly bool,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	model_context.New(s, readOnly, r, t, version).Mount(g)
}
