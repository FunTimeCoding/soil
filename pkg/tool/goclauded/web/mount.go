package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc(
		route.Get(webConstant.PalettePath),
		palette.NewServe(s.registry),
	)
	m.HandleFunc(route.Get(webConstant.RootPattern), s.dashboard)
	m.Handle(route.Get(webConstant.LivePath), s.event())
	m.HandleFunc(route.Get(constant.SessionsPath), s.sessionsPage)
	m.HandleFunc(
		route.Get(constant.SessionsPath, "/{identifier}"),
		s.sessionDetailPage,
	)
	m.HandleFunc(
		route.Get(constant.SessionsPath, "/{identifier}/edit"),
		s.sessionEditForm,
	)
	m.HandleFunc(
		route.Post(constant.SessionsPath, "/{identifier}/edit"),
		s.sessionEditSubmit,
	)
	m.HandleFunc(
		route.Post(constant.SessionsPath, "/{identifier}/pulse"),
		s.sessionPulseSubmit,
	)
	m.HandleFunc(
		route.Post(constant.SessionsPath, "/{identifier}/delete"),
		s.sessionDeleteAction,
	)
	m.HandleFunc(route.Get("/activity"), s.activityPage)
	m.HandleFunc(route.Get(constant.MessagesPath), s.messagesPage)
	m.HandleFunc(route.Get(constant.HistoryPath), s.historyPage)
	m.HandleFunc(
		route.Get(constant.HistoryPath, "/{identifier}/edit"),
		s.historyEditForm,
	)
	m.HandleFunc(
		route.Post(constant.HistoryPath, "/{identifier}/edit"),
		s.historyEditSubmit,
	)
	m.HandleFunc(route.Get(webConstant.FaviconPath), s.favicon)
	s.conversations.Mount(m)
}
