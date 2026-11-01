package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/status_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) statusPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	findings, e := s.service.Findings()

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderPage(
		w,
		constant.StatusTitle,
		constant.StatusPath,
		html.H3(gomponents.Text(constant.StatusTitle)),
		status_page.Table(findings),
	)
}
