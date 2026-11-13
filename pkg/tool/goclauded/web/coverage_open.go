package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/coverage_page"
	"net/http"
)

func (s *Server) coverageOpen(
	w http.ResponseWriter,
	q *http.Request,
) {
	name := q.URL.Query().Get(constant.CoverageServer)
	e := coverageServer(s.transcript.Coverage(), name)

	if e == nil {
		s.view.RenderFragment(w, coverage_page.Removal(name))

		return
	}

	s.view.RenderFragment(w, coverage_page.Detail(e))
}
