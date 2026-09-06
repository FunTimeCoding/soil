package godashboardd

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	u *web.Server,
	g *guard.Mux,
) {
	u.Mount(g)
}
