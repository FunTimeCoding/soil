package web

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/flag_page"
	"net/http"
)

func (s *Server) flagOpen(
	w http.ResponseWriter,
	r *http.Request,
) {
	row, e := s.overview.Flagged(
		r.URL.Query().Get(generativeConstant.ParameterIdentifier),
	)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	s.view.RenderFragment(w, flag_page.Detail(row))
}
