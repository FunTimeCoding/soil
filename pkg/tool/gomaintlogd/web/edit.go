package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web/entry_page"
	"net/http"
)

func (s *Server) edit(
	w http.ResponseWriter,
	r *http.Request,
) {
	e, f := s.entryFromQuery(r)

	if f != nil {
		s.view.Answer(w, r, s.reporter, f)

		return
	}

	s.view.RenderFragment(w, entry_page.EditRow(e, ""))
}
