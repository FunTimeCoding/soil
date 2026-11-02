package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/http"
)

func (s *Server) historyEditSubmit(
	w http.ResponseWriter,
	q *http.Request,
) {
	e, okay := s.historyEvent(w, q)

	if !okay {
		return
	}

	s.event.MustEditEvent(e.Identifier, q.FormValue(constant.Body))
	s.historyRow(w, q)
}
