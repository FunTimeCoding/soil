package gofirefoxd

import (
	"github.com/funtimecoding/soil/pkg/face"
	firefox "github.com/funtimecoding/soil/pkg/tool/gofirefoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	c firefox.FirefoxSource,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	model_context.New(c, r, t, version).Mount(g)
}
