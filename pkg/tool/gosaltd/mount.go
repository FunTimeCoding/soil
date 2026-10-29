package gosaltd

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/provision"
	salt "github.com/funtimecoding/soil/pkg/tool/gosaltd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/model_context"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	v salt.Salt,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	provision.Mount(v, r, t, g)
	model_context.New(v, r, t).Mount(g)
}
