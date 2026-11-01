package web

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"
	"net/http"
	"strconv"
)

func (s *Server) historyEvent(
	w http.ResponseWriter,
	q *http.Request,
) (*event.Event, bool) {
	raw := q.PathValue(constant.Identifier)
	identifier, e := strconv.ParseUint(raw, 10, 64)
	var found *event.Event

	if e == nil && identifier > 0 {
		found = s.service.GetEvent(uint(identifier))
	}

	if found == nil {
		s.view.Answer(w, q, s.reporter, not_found.New(constant.EventKind, raw))

		return nil, false
	}

	return found, true
}
