package web

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	item := "/{identifier}"
	edit := join.Empty(item, stringsConstant.Slash, constant.EditAction)
	pulse := join.Empty(item, stringsConstant.Slash, constant.PulseAction)
	cruise := join.Empty(item, stringsConstant.Slash, constant.CruiseAction)
	remove := join.Empty(item, stringsConstant.Slash, constant.DeleteAction)
	g.Open(route.Get(webConstant.PalettePath), s.registry.Serve())
	g.Open(route.Get(webConstant.RootPattern), s.dashboard)
	g.OpenMount(route.Get(webConstant.LivePath), s.event())
	g.Token(route.Get(constant.EventStreamPath), s.eventStream())
	g.Open(route.Get(constant.SessionsPath), s.sessionsPage)
	g.Open(route.Get(constant.SessionsPath, item), s.sessionDetailPage)
	g.Open(route.Get(constant.SessionsPath, edit), s.sessionEditForm)
	g.Open(route.Post(constant.SessionsPath, edit), s.sessionEditSubmit)
	g.Open(route.Post(constant.SessionsPath, pulse), s.sessionPulseSubmit)
	g.Open(route.Post(constant.SessionsPath, cruise), s.sessionCruiseSubmit)
	g.Open(route.Post(constant.SessionsPath, remove), s.sessionDeleteAction)
	g.Open(route.Get(constant.ActivityPath), s.activityPage)
	g.Open(route.Get(constant.CoveragePath), s.coveragePage)
	g.Open(route.Get(constant.StatusPath), s.statusPage)
	g.Open(route.Get(constant.UsagePath), s.usagePage)
	g.Open(route.Get(constant.MessagesPath), s.messagesPage)
	g.Open(route.Get(constant.HistoryPath), s.historyPage)
	g.Open(route.Get(constant.HistoryPath, item), s.historyRow)
	g.Open(route.Get(constant.HistoryPath, edit), s.historyEditForm)
	g.Open(route.Post(constant.HistoryPath, edit), s.historyEditSubmit)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
	s.conversations.Mount(g)
}
