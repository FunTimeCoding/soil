package gosentryd

import (
	"github.com/funtimecoding/soil/pkg/face"
	sentry "github.com/funtimecoding/soil/pkg/tool/gosentryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	c sentry.SentrySource,
	organization string,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	model_context.New(c, organization, r, t, version).Mount(g)
}
