package goitermd

import (
	"github.com/funtimecoding/soil/pkg/face"
	iterm "github.com/funtimecoding/soil/pkg/tool/goitermd/face"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	v iterm.Iterm,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	model_context.New(v, r, t).Mount(g)
}
