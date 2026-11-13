package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/conversation_page"
	"net/http"
	"strconv"
)

func (s *Server) conversationSidebar(
	w http.ResponseWriter,
	q *http.Request,
) {
	skip := 0

	if v := q.URL.Query().Get(constant.Offset); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			skip = n
		}
	}

	sessions, next, e := s.conversationSessions(skip)

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderFragment(w, conversation_page.Sidebar(sessions, "", next))
}
