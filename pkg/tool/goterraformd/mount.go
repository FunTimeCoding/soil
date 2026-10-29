package goterraformd

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/provision"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/model_context"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	v provisionFace.Provision,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	provision.Mount(v, r, t, g)
	model_context.New(v, r, t).Mount(g)
}
