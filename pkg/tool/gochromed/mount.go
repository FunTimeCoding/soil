package gochromed

import (
	"github.com/funtimecoding/soil/pkg/face"
	chromium "github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	v chromium.Browser,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	model_context.New(v, r, t).Mount(g)
}
