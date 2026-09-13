package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) notification(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderPage(
		w,
		constant.NotifyTitle,
		constant.NotifyPath,
		html.H1(gomponents.Text(constant.NotifyTitle)),
		html.Form(
			html.ID("receipt-form"),
			extended.Post(constant.OutOfBandPath),
			extended.Target(selector(constant.ReceiptMark)),
			extended.Swap(constant.SwapInner),
			extended.AfterRequest("this.reset()"),
			html.Input(
				html.ID("note-input"),
				html.Type("text"),
				html.Name(constant.TermField),
			),
			html.Button(html.ID("receipt-post"), gomponents.Text("Send")),
		),
		html.Div(html.ID(constant.ReceiptMark)),
		html.Div(html.ID(constant.SummaryMark)),
		html.Button(
			html.ID("fail-post"),
			extended.Post(constant.FailPath),
			extended.Target(selector(constant.ReceiptMark)),
			extended.Swap(constant.SwapInner),
			gomponents.Text("Fail"),
		),
	)
}
