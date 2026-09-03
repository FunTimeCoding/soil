package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(constant.PalettePath), palette.NewServe(s.registry))
	g.Open(route.Get(constant.RootPattern), s.dashboard)
	g.OpenMount(route.Get(constant.LivePath), s.event())
	g.Open(route.Post("/move-up"), s.moveUp)
	g.Open(route.Post("/move-down"), s.moveDown)
	g.Open(route.Post("/reorder"), s.reorder)
	g.Open(route.Get(constant.FaviconPath), s.favicon)
}
