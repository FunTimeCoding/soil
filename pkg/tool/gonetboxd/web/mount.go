package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(constant.PalettePath), s.registry.Serve())
	g.Open(route.Get(constant.RootPattern), s.bookmarks)
	g.Open(route.Get(constant.FaviconPath), s.favicon)
}
