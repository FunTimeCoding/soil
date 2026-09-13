package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) stream(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderLivePage(
		w,
		constant.StreamTitle,
		constant.StreamPath,
		subscription.Query(constant.TickMark, constant.PulseMark),
		html.H1(gomponents.Text(constant.StreamTitle)),
		html.Div(
			html.ID(constant.TickMark),
			extended.StreamSwap(constant.TickMark),
			s.tickCell(),
		),
		html.Div(
			html.ID(constant.PulseMark),
			extended.StreamSwap(constant.PulseMark),
			s.pulseCell(),
		),
		html.Button(
			html.ID("pulse-post"),
			extended.Post(constant.SubscribePath),
			extended.Target(selector(constant.ReceiptMark)),
			extended.Swap(constant.SwapInner),
			gomponents.Text("Pulse"),
		),
		html.Div(html.ID(constant.ReceiptMark)),
	)
}
