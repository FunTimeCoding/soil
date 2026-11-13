package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/conversation_page"
	"net/http"
)

func (s *Server) conversationPanel(
	w http.ResponseWriter,
	q *http.Request,
) {
	identifier := q.PathValue(constant.Identifier)
	messages, e := s.conversationMessages(identifier)

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	panel := conversation_page.Messages(messages)

	if s.view.IsExtendedRequest(q) {
		s.view.RenderFragment(w, panel)

		return
	}

	s.renderConversations(w, q, identifier, panel)
}
