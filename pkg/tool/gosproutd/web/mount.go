package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(webConstant.PalettePath), s.registry.Serve())
	g.Open(route.Get(webConstant.RootPattern), s.dashboard)
	g.OpenMount(route.Get(webConstant.LivePath), s.event())
	g.Open(route.Post(constant.ReorderPath), s.reorder)
	g.Open(route.Get(constant.SessionsPath), s.sessions)
	g.Open(route.Get(constant.CountersPath), s.counters)
	g.Open(route.Get(constant.SessionPath), s.sessionDetail)
	g.Open(route.Get(constant.DecisionPath), s.decision)
	g.Open(route.Post(constant.AnswerPath), s.answer)
	g.Open(route.Post(constant.DismissPath), s.dismiss)
	g.Open(route.Post(constant.ReplyPath), s.reply)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
}
