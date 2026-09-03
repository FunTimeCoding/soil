package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(constant.PalettePath), palette.NewServe(s.registry))
	g.Open(route.Get(constant.RootPattern), s.floor)
	g.OpenMount(route.Get(constant.LivePath), s.event())
	g.Open(route.Get(constant.FaviconPath), s.favicon)
}
