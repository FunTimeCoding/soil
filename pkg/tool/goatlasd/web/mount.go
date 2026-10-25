package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	s.authorization.Mount(g)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
	g.Session(route.Get(webConstant.PalettePath), s.palette.Serve())
	g.Session(route.Get(webConstant.RootPattern), s.dashboard)
	g.Session(route.Get(constant.PlacementsPath), s.placements)
	g.Session(route.Get(constant.SightingsPath), s.sightings)
	g.Session(
		route.Get(
			constant.PlacesPath,
			"/{",
			constant.KindParameter,
			"}/{",
			constant.NameParameter,
			"}",
		),
		s.placeDetail,
	)
}
