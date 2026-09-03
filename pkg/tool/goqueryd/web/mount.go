package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(webConstant.PalettePath), palette.NewServe(s.registry))
	g.Open(route.Get(webConstant.RootPattern), s.dashboard)
	g.Open(route.Get(webConstant.SearchPath), s.searchPage)
	g.Open(route.Get(constant.CollectionsPath), s.collectionsPage)
	g.Open(route.Get(constant.CollectionsPath, "/{name}"), s.documentsPage)
	g.Open(route.Get("/documents/{path...}"), s.documentDetailPage)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
