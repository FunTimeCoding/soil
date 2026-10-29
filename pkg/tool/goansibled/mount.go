package goansibled

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/provision"
	ansible "github.com/funtimecoding/soil/pkg/tool/goansibled/face"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/model_context"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	v ansible.Ansible,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	provision.Mount(v, r, t, g)
	model_context.New(v, r, t).Mount(g)
}
