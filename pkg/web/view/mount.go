package view

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (v *View) Mount(g *guard.Mux) {
	g.Open(route.Get(constant.FaviconPath), v.Favicon)

	if g.HasSession() {
		g.Session(route.Get(constant.PalettePath), v.Palette)

		return
	}

	g.Open(route.Get(constant.PalettePath), v.Palette)
}
