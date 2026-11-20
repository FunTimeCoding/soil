package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"net/url"
)

func (s *Server) flags(
	w http.ResponseWriter,
	r *http.Request,
) {
	section, e := s.section(
		url.Values{},
		"",
		r.URL.Query().Get(constant.ShowField),
	)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(w, section)

		return
	}

	s.view.RenderPage(
		w,
		constant.FlagsTitle,
		constant.FlagsPath,
		html.H1(gomponents.Text(constant.FlagsTitle)),
		section,
	)
}
