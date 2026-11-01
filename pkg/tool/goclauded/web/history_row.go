package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/timeline"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/history_page"
	"net/http"
)

func (s *Server) historyRow(
	w http.ResponseWriter,
	q *http.Request,
) {
	e, okay := s.historyEvent(w, q)

	if !okay {
		return
	}

	entry := timeline.FromEvent(
		e.Identifier,
		e.SessionIdentifier,
		e.Kind,
		e.Actor,
		e.Metadata,
		e.CreatedAt,
	)
	entry.Full = true
	s.view.RenderFragment(w, history_page.Row(entry, true))
}
