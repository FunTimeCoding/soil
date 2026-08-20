package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc(route.Get(webConstant.SignInPath), s.signIn)
	m.HandleFunc(route.Get(webConstant.CallbackPath), s.callback)
	m.HandleFunc(route.Get(webConstant.SignOutPath), s.signOut)
	m.HandleFunc(
		route.Get(webConstant.PalettePath),
		s.require(palette.NewServe(s.registry)),
	)
	m.HandleFunc(route.Get(webConstant.RootPattern), s.require(s.dashboard))
	m.HandleFunc(route.Get(constant.HeatmapPath), s.require(s.heatmap))
	m.HandleFunc(route.Post("/click"), s.require(s.click))
	m.Handle(route.Get(webConstant.LivePath), s.require(s.event()))
	m.HandleFunc(route.Get(webConstant.FaviconPath), s.favicon)
}
