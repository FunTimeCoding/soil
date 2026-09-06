package goprocessd

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/server"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	s *server.Server,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	model_context.New(s, r, t, version).Mount(g)
}
