package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func (s *Server) Mount(g *guard.Mux) {
	s.view.Mount(g)
	g.Session(route.Get(webConstant.RootPattern), s.logs)
	g.Session(route.Post(constant.GeneratePath), s.generate)
	g.Session(route.Get(constant.ReportsPath), s.reports)
	g.Session(route.Get(constant.ReportsPath, "/{fileName}"), s.reportDownload)
	g.Session(
		route.Post(constant.ReportsPath, "/{fileName}", constant.DeleteSuffix),
		s.reportDelete,
	)
	g.Session(route.Get(constant.RaidsPath), s.raids)
	g.Session(route.Get(constant.RaidsPath, "/{id}"), s.raidDetail)
	g.Session(route.Post(constant.CreateRaidPath), s.createRaid)
	g.Session(route.Get(constant.PlayersPath), s.players)
	g.Session(route.Get(constant.PlayersPath, "/{account}"), s.playerDetail)
	g.OpenMount(route.Get(constant.StaticPath), http.FileServerFS(staticFiles))
}
