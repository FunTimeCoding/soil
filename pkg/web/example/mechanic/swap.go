package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) swap(
	w http.ResponseWriter,
	r *http.Request,
) {
	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(w, s.counterCell())

		return
	}

	s.view.RenderPage(
		w,
		constant.SwapTitle,
		constant.RootPath,
		html.H1(gomponents.Text(constant.SwapTitle)),
		html.Div(html.ID(constant.CounterMark), s.counterCell()),
		html.Button(
			html.ID("counter-get"),
			extended.Get(constant.CounterPath),
			extended.Target(selector(constant.CounterMark)),
			extended.Swap(constant.SwapInner),
			gomponents.Text("Read"),
		),
		html.Button(
			html.ID("counter-post"),
			extended.Post(constant.CounterPath),
			extended.Target(selector(constant.CounterMark)),
			extended.Swap(constant.SwapInner),
			gomponents.Text("Increment"),
		),
		rowCell(false),
		html.Button(
			html.ID("row-post"),
			extended.Post(constant.RowPath),
			extended.Target(selector(constant.RowMark)),
			extended.Swap(constant.SwapOuter),
			extended.Confirm(constant.ConfirmMessage),
			gomponents.Text("Replace row"),
		),
	)
}
