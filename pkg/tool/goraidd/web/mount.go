package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
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
	m.HandleFunc(route.Get(webConstant.RootPattern), s.require(s.logs))
	m.HandleFunc(route.Post("/generate"), s.require(s.generate))
	m.HandleFunc(route.Get(constant.ReportsPath), s.require(s.reports))
	m.HandleFunc(
		route.Get(constant.ReportsPath, "/{fileName}"),
		s.require(s.reportDownload),
	)
	m.HandleFunc(
		route.Post(constant.ReportsPath, "/{fileName}/delete"),
		s.require(s.reportDelete),
	)
	m.HandleFunc(route.Get(constant.RaidsPath), s.require(s.raids))
	m.HandleFunc(
		route.Get(constant.RaidsPath, "/{id}"),
		s.require(s.raidDetail),
	)
	m.HandleFunc(
		route.Post(constant.RaidsPath, "/create"),
		s.require(s.createRaid),
	)
	m.HandleFunc(route.Get(constant.PlayersPath), s.require(s.players))
	m.HandleFunc(
		route.Get(constant.PlayersPath, "/{account}"),
		s.require(s.playerDetail),
	)
	m.HandleFunc(route.Get(webConstant.FaviconPath), s.favicon)
	m.Handle(route.Get("/static/"), http.FileServerFS(staticFiles))
}
