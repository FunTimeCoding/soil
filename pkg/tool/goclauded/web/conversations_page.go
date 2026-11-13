package web

import "net/http"

func (s *Server) conversationsPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	s.renderConversations(w, q, "", nil)
}
