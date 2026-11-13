package web

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	s.view.Mount(g)
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
	item := "/{identifier}"
	open := join.Empty(item, stringsConstant.Slash, constant.OpenAction)
	closing := join.Empty(item, stringsConstant.Slash, constant.CloseAction)
	edit := join.Empty(item, stringsConstant.Slash, constant.EditAction)
	g.Open(route.Get(constant.MemoriesPath, item), s.memoryDetailPage)
	g.Open(route.Get(constant.MemoriesPath, open), s.memoryOpen)
	g.Open(route.Get(constant.MemoriesPath, closing), s.memoryClose)
	g.Open(route.Get(constant.MemoriesPath, edit), s.memoryEdit)
	g.Open(route.Post(constant.MemoriesPath, edit), s.memoryEditSubmit)
	g.Open(route.Get(constant.RelationsPath), s.relationsPage)
	g.Open(route.Get(constant.StatisticPath), s.statisticPage)
	g.Open(route.Get(constant.ImpressionsPath), s.impressionsPage)
	g.Open(route.Get(webConstant.SearchPath), s.searchPage)
}
