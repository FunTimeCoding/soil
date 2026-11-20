package web

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/url"
)

func (s *Server) renderSection(
	w http.ResponseWriter,
	r *http.Request,
	values url.Values,
	alert string,
	success string,
) {
	section, e := s.section(values, alert, "")

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	if success == "" {
		s.view.RenderFragment(w, section)

		return
	}

	s.view.RenderStatusFragment(
		w,
		section,
		success,
		webConstant.LayoutStatusSuccess,
	)
}
