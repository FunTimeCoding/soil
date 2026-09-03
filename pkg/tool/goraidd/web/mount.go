package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func (s *Server) Mount(g *guard.Mux) {
	g.WithSession(s.require)
	g.Open(route.Get(webConstant.SignInPath), s.signIn)
	g.Open(route.Get(webConstant.CallbackPath), s.callback)
	g.Open(route.Get(webConstant.SignOutPath), s.signOut)
	g.Session(route.Get(webConstant.PalettePath), palette.NewServe(s.registry))
	g.Session(route.Get(webConstant.RootPattern), s.logs)
	g.Session(route.Post("/generate"), s.generate)
	g.Session(route.Get(constant.ReportsPath), s.reports)
	g.Session(route.Get(constant.ReportsPath, "/{fileName}"), s.reportDownload)
	g.Session(
		route.Post(constant.ReportsPath, "/{fileName}/delete"),
		s.reportDelete,
	)
	g.Session(route.Get(constant.RaidsPath), s.raids)
	g.Session(route.Get(constant.RaidsPath, "/{id}"), s.raidDetail)
	g.Session(route.Post(constant.RaidsPath, "/create"), s.createRaid)
	g.Session(route.Get(constant.PlayersPath), s.players)
	g.Session(route.Get(constant.PlayersPath, "/{account}"), s.playerDetail)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
	g.OpenMount(route.Get("/static/"), http.FileServerFS(staticFiles))
}
