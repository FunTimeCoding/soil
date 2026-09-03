package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(webConstant.PalettePath), palette.NewServe(s.registry))
	g.Open(route.Get(webConstant.RootPattern), s.dashboard)
	g.Open(route.Get(constant.RecentPath), s.recent)
	g.Open(route.Get("/alerts"), s.alerts)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
