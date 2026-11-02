package web

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(webConstant.PalettePath), s.registry.Serve())
	g.Open(route.Get(webConstant.RootPattern), s.bookmarks)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
