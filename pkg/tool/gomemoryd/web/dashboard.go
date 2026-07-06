package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderPage(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		s.memorySummary(),
		s.tagSummary(),
		s.recentVersions(),
		s.recentImpressionsSummary(),
	)
}
