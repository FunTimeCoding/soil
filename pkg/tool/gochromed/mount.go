package gochromed

import (
	"github.com/funtimecoding/soil/pkg/face"
	chromium "github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	c chromium.ChromiumSource,
	downloadDirectory string,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	model_context.New(c, downloadDirectory, r, t, version).Mount(g)
}
