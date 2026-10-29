package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	s.service.Authorization().Mount(g)
	g.Session(route.Get(webConstant.PalettePath), s.registry.Serve())
	g.Session(route.Get(webConstant.RootPattern), s.dashboard)
	g.Session(route.Get(constant.HeatmapPath), s.heatmap)
	g.Session(route.Post("/click"), s.click)
	g.Session(route.Get(webConstant.LivePath), s.event())
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
