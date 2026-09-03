package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(webConstant.PalettePath), palette.NewServe(s.registry))
	g.Open(
		route.Get(webConstant.PalettePath, constant.MemoriesPath),
		s.paletteMemories,
	)
	g.Open(
		route.Get(
			webConstant.PalettePath,
			constant.MemoriesPath,
			webConstant.SearchPath,
		),
		s.paletteMemoriesSearch,
	)
	g.Open(route.Get(webConstant.RootPattern), s.dashboard)
	g.Open(route.Get(constant.MemoriesPath), s.memoriesPage)
	g.Open(
		route.Get(constant.MemoriesPath, "/{identifier}"),
		s.memoryDetailPage,
	)
	g.Open(route.Get(constant.RelationsPath), s.relationsPage)
	g.Open(route.Get(constant.ImpressionsPath), s.impressionsPage)
	g.Open(route.Get(webConstant.SearchPath), s.searchPage)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
