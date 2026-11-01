package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/counter_page"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) counters(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderPage(
		w,
		constant.CountersTitle,
		constant.CountersPath,
		html.H3(gomponents.Text(constant.CountersTitle)),
		counter_page.Note(),
		html.Div(
			extended.StreamSwap(constant.CounterEvent),
			counter_page.Table(s.service.Tallies()),
		),
	)
}
