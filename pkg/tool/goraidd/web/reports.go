package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) reports(
	w http.ResponseWriter,
	r *http.Request,
) {
	reports, e := s.service.Reports()
	errors.PanicOnError(e)

	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(w, reportsTable(reports))

		return
	}

	s.view.RenderPage(
		w,
		constant.ReportsTitle,
		constant.ReportsPath,
		html.H1(gomponents.Text("Generated Reports")),
		reportsTable(reports),
	)
}
