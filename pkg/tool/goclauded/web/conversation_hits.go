package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/conversation_page"
	"net/http"
)

func (s *Server) conversationHits(
	w http.ResponseWriter,
	q *http.Request,
) {
	session := q.PathValue(constant.Identifier)
	query := q.URL.Query().Get(constant.QueryField)
	kinds := searchKinds(q)
	blocks, e := s.search.ConversationBlocks(session, kinds)

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	matching, f := s.search.MatchingBlocks(session, query, kinds)

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	panel := conversation_page.Hits(
		blocks,
		matching,
		query,
		q.URL.Query().Get(constant.Around),
	)

	if s.view.IsExtendedRequest(q) {
		s.view.RenderFragment(w, panel)

		return
	}

	s.renderConversations(w, q, "", panel)
}
