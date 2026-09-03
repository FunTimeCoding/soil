package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (s *Server) Mount(g *guard.Mux) {
	g.Open(route.Get(webConstant.PalettePath), palette.NewServe(s.registry))
	g.Open(route.Get(webConstant.RootPattern), s.dashboard)
	g.OpenMount(route.Get(webConstant.LivePath), s.event())
	g.Open(route.Get(constant.SessionsPath), s.sessionsPage)
	g.Open(
		route.Get(constant.SessionsPath, "/{identifier}"),
		s.sessionDetailPage,
	)
	g.Open(
		route.Get(constant.SessionsPath, "/{identifier}/edit"),
		s.sessionEditForm,
	)
	g.Open(
		route.Post(constant.SessionsPath, "/{identifier}/edit"),
		s.sessionEditSubmit,
	)
	g.Open(
		route.Post(constant.SessionsPath, "/{identifier}/pulse"),
		s.sessionPulseSubmit,
	)
	g.Open(
		route.Post(constant.SessionsPath, "/{identifier}/delete"),
		s.sessionDeleteAction,
	)
	g.Open(route.Get("/activity"), s.activityPage)
	g.Open(route.Get(constant.CoveragePath), s.coveragePage)
	g.Open(route.Get(constant.UsagePath), s.usagePage)
	g.Open(route.Get(constant.MessagesPath), s.messagesPage)
	g.Open(route.Get(constant.HistoryPath), s.historyPage)
	g.Open(
		route.Get(constant.HistoryPath, "/{identifier}/edit"),
		s.historyEditForm,
	)
	g.Open(
		route.Post(constant.HistoryPath, "/{identifier}/edit"),
		s.historyEditSubmit,
	)
	g.Open(route.Get(webConstant.FaviconPath), s.favicon)
	s.conversations.Mount(g)
}
