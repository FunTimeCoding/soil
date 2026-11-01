package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/dashboard_page"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	q *http.Request,
) {
	v, e := s.service.DashboardView()

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	p := dashboard_page.New(v)
	s.view.RenderPageWithSummary(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		p.Summary(),
		p.Memories(),
		p.Tags(),
		p.Versions(),
		p.Impressions(),
	)
}
