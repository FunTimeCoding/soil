package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/history_page"
	"net/http"
)

func (s *Server) historyEditForm(
	w http.ResponseWriter,
	q *http.Request,
) {
	e, okay := s.historyEvent(w, q)

	if !okay {
		return
	}

	s.view.RenderFragment(w, history_page.EditRow(e))
}
