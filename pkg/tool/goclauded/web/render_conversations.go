package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/conversation_page"
	"maragu.dev/gomponents"
	"net/http"
)

func (s *Server) renderConversations(
	w http.ResponseWriter,
	q *http.Request,
	current string,
	panel gomponents.Node,
) {
	query := q.URL.Query().Get(constant.QueryField)
	kinds := searchKinds(q)
	sidebar, e := s.conversationSidebarNode(query, kinds, current)

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderFullWindowPage(
		w,
		constant.ConversationsTitle,
		constant.ConversationsPath,
		conversation_page.Layout(query, kinds, sidebar, panel),
	)
}
