package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(webConstant.PalettePath), s.registry.Serve())
	g.Open(route.Get(webConstant.RootPattern), s.heatmap)
	g.Open(route.Get(constant.EventsPath), s.events)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
