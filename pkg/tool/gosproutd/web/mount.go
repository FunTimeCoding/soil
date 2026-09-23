package web

import (
	sproutConstant "github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(constant.PalettePath), palette.NewServe(s.registry))
	g.Open(route.Get(constant.RootPattern), s.dashboard)
	g.OpenMount(route.Get(constant.LivePath), s.event())
	g.Open(route.Post("/move-up"), s.moveUp)
	g.Open(route.Post("/move-down"), s.moveDown)
	g.Open(route.Post("/reorder"), s.reorder)
	g.Open(route.Get(sproutConstant.SessionsPath), s.sessions)
	g.Open(route.Get(sproutConstant.CountersPath), s.counters)
	g.Open(route.Get(sproutConstant.SessionPath), s.sessionDetail)
	g.Open(route.Get(sproutConstant.DecisionPath), s.decision)
	g.Open(route.Post(sproutConstant.AnswerPath), s.answer)
	g.Open(route.Post(sproutConstant.DismissPath), s.dismiss)
	g.Open(route.Post(sproutConstant.ReplyPath), s.reply)
	g.Open(route.Get(constant.FaviconPath), s.favicon)
}
