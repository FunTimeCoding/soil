package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/http"
)

func (s *Server) conversationSearch(
	w http.ResponseWriter,
	q *http.Request,
) {
	if !s.view.IsExtendedRequest(q) {
		s.renderConversations(w, q, "", nil)

		return
	}

	sidebar, e := s.conversationSidebarNode(
		q.URL.Query().Get(constant.QueryField),
		searchKinds(q),
		"",
	)

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderFragment(w, sidebar)
}
