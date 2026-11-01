package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/statistic_page"
	"net/http"
)

func (s *Server) statisticPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	summary, e := s.service.MemoryTokenSummary(
		q.URL.Query().Get(constant.Scope),
	)

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderPage(
		w,
		constant.StatisticTitle,
		constant.StatisticPath,
		statistic_page.Region(summary),
	)
}
