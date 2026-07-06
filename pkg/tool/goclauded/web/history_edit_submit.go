package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/timeline"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"strconv"
)

func (s *Server) historyEditSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier, e := strconv.Atoi(r.PathValue(constant.Identifier))

	if e != nil || identifier <= 0 {
		return
	}

	errors.PanicOnError(r.ParseForm())
	body := r.FormValue(constant.Body)
	s.service.MustEditEvent(uint(identifier), body)
	event := s.service.GetEvent(uint(identifier))

	if event == nil {
		return
	}

	entry := timeline.FromEvent(
		event.Identifier,
		event.SessionIdentifier,
		event.Kind,
		event.Actor,
		event.Metadata,
		event.CreatedAt,
	)
	entry.Full = true
	w.Header().Set(web.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(timelineRow(entry).Render(w))
}
