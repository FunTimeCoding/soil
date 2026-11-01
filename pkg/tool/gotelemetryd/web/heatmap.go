package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/summary_option"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/web/summary_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) heatmap(
	w http.ResponseWriter,
	r *http.Request,
) {
	q := r.URL.Query()
	window := q.Get(constant.WindowParameter)

	if window == "" {
		window = constant.DefaultWindow
	}

	group := q.Get(constant.GroupParameter)

	if group == "" {
		group = constant.Tool
	}

	o := summary_option.New()
	o.Since = sinceFromWindow(window)
	o.GroupBy = group
	rows, e := s.service.Summary(o)
	errors.PanicOnError(e)
	s.view.RenderPage(
		w,
		"",
		constant.HeatmapPath,
		html.H3(gomponents.Text(constant.HeatmapTitle)),
		summary_page.Selectors(window, group),
		summary_page.Table(rows, group),
	)
}
