package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"net/http"
)

func (s *Server) heatmap(
	w http.ResponseWriter,
	_ *http.Request,
) {
	summaries, e := s.store.Summaries()
	errors.PanicOnError(e)
	s.view.RenderPage(
		w,
		constant.HeatmapTitle,
		constant.HeatmapPath,
		heatmapTable(summaries),
	)
}
