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

	since := time.Now().Add(-time.Duration(days) * 24 * time.Hour).UTC()
	impressions, f := s.service.RecentImpressions(since.Format(time.RFC3339))

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
