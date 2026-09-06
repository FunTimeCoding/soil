package goitermd

import (
	"github.com/funtimecoding/soil/pkg/face"
	iterm "github.com/funtimecoding/soil/pkg/tool/goitermd/face"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	c iterm.ItermSource,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	model_context.New(c, r, t, version).Mount(g)
}
