package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
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
		html.P(
			html.Small(
				html.Style("color: var(--pico-muted-color);"),
				gomponents.Text(
					"Irrelevant counts my bad asks. Abandoned counts the ones you never needed to answer. Both are measurements of me, not of you.",
				),
			),
		),
		html.Div(extended.StreamSwap(constant.CounterEvent), s.tallyTable()),
	)
}
