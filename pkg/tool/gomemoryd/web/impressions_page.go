package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/impression_page"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) impressionsPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	days, e := strconv.Atoi(q.URL.Query().Get(constant.Days))

	if e != nil || days < 1 {
		days = constant.DefaultImpressionDays
	}

	impressions, f := s.service.RecentImpressions(
		time.Duration(days) * 24 * time.Hour,
	)

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	s.view.RenderPage(
		w,
		constant.ImpressionsTitle,
		constant.ImpressionsPath,
		impression_page.List(impressions, days),
	)
}
