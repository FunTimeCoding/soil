package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/surface"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"time"
)

func (d *Daemon) Mount(
	r face.Reporter,
	hold time.Duration,
	t face.Recorder,
	g *guard.Mux,
) {
	surface.Mount(d.service, d.web, d.logger, r, hold, t, g)
}
