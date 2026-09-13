package mechanic

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) extra(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderPage(
		w,
		constant.ExtraTitle,
		constant.ExtraPath,
		html.H1(gomponents.Text(constant.ExtraTitle)),
		html.Div(html.ID(constant.RemoveMark), gomponents.Text("removable")),
		html.Button(
			html.ID("remove-post"),
			extended.Post(constant.RemovePath),
			extended.Target(selector(constant.RemoveMark)),
			extended.Swap(constant.SwapDelete),
			gomponents.Text("Remove"),
		),
		html.Ul(html.ID(constant.AppendMark)),
		html.Button(
			html.ID("append-post"),
			extended.Post(constant.AppendPath),
			extended.Target(selector(constant.AppendMark)),
			extended.Swap(constant.SwapAppend),
			gomponents.Text("Append"),
		),
		html.Div(html.ID(constant.QuietMark), gomponents.Text("untouched")),
		html.Button(
			html.ID("quiet-post"),
			extended.Post(constant.QuietPath),
			extended.Target(selector(constant.QuietMark)),
			extended.Swap(constant.SwapNone),
			gomponents.Text("Quiet"),
		),
		html.Div(
			html.ID(constant.BranchMark),
			extended.Get(constant.BranchPath),
			extended.Trigger(constant.TriggerLoad),
			extended.Swap(constant.SwapInner),
		),
		html.Div(
			html.ID(constant.PollMark),
			extended.Get(constant.PollPath),
			extended.Trigger(constant.PollWhen),
			extended.Swap(constant.SwapInner),
		),
		html.Div(html.ID(constant.FireMark)),
		layout.Indicator(constant.IndicatorMark, constant.IndicatorText),
		html.Div(html.ID(constant.SlowMark)),
		html.Button(
			html.ID(constant.SlowControlMark),
			extended.Post(constant.SlowPath),
			extended.Target(selector(constant.SlowMark)),
			extended.Swap(constant.SwapInner),
			extended.Indicator(
				join.Comma(
					[]string{
						selector(constant.SlowControlMark),
						selector(constant.IndicatorMark),
					},
				),
			),
			gomponents.Text("Slow"),
		),
		html.Button(
			html.ID("redirect-post"),
			extended.Post(constant.RedirectPath),
			extended.Target(selector(constant.QuietMark)),
			extended.Swap(constant.SwapNone),
			gomponents.Text("Redirect"),
		),
		html.Button(
			html.ID("refresh-post"),
			extended.Post(constant.RefreshPath),
			extended.Target(selector(constant.QuietMark)),
			extended.Swap(constant.SwapNone),
			gomponents.Text("Refresh"),
		),
		html.Button(
			html.ID("fire-post"),
			extended.Post(constant.FirePath),
			extended.Target(selector(constant.QuietMark)),
			extended.Swap(constant.SwapNone),
			gomponents.Attr(
				join.Empty(constant.ExtendedOnPrefix, constant.FiredEvent),
				"document.getElementById('fire').textContent = 'heard'",
			),
			gomponents.Text("Fire"),
		),
	)
}
