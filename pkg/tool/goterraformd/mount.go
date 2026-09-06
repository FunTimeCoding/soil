package goterraformd

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/runner"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	n *runner.Runner,
	s *store.Store,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	model_context.New(n, s, r, t, version).Mount(g)
}
