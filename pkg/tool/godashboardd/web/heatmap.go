package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/web/heatmap_page"
	"net/http"
)

func (s *Server) heatmap(
	w http.ResponseWriter,
	r *http.Request,
) {
	summaries, e := s.service.Summaries()

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	s.view.RenderPage(
		w,
		constant.HeatmapTitle,
		constant.HeatmapPath,
		heatmap_page.Table(summaries),
	)
}
