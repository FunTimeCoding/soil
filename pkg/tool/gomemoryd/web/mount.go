package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc(
		route.Get(webConstant.PalettePath),
		palette.NewServe(s.registry),
	)
	m.HandleFunc(
		route.Get(webConstant.PalettePath, constant.MemoriesPath),
		s.paletteMemories,
	)
	m.HandleFunc(
		route.Get(
			webConstant.PalettePath,
			constant.MemoriesPath,
			webConstant.SearchPath,
		),
		s.paletteMemoriesSearch,
	)
	m.HandleFunc(route.Get(webConstant.RootPattern), s.dashboard)
	m.HandleFunc(route.Get(constant.MemoriesPath), s.memoriesPage)
	m.HandleFunc(
		route.Get(constant.MemoriesPath, "/{identifier}"),
		s.memoryDetailPage,
	)
	m.HandleFunc(route.Get(constant.RelationsPath), s.relationsPage)
	m.HandleFunc(route.Get(constant.ImpressionsPath), s.impressionsPage)
	m.HandleFunc(route.Get(webConstant.SearchPath), s.searchPage)
	m.HandleFunc(route.Get(webConstant.FaviconPath), s.favicon)
}
