package mechanic

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) trigger(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderPage(
		w,
		constant.TriggerTitle,
		constant.TriggerPath,
		html.H1(gomponents.Text(constant.TriggerTitle)),
		html.Div(
			html.ID(constant.LoadMark),
			extended.Get(join.Empty(constant.EchoPath, constant.LoadQuery)),
			extended.Trigger(constant.TriggerLoad),
			extended.Swap(constant.SwapInner),
		),
		html.Div(
			html.ID(constant.DelayedMark),
			extended.Get(join.Empty(constant.EchoPath, constant.DelayedQuery)),
			extended.Trigger(constant.DelayWhen),
			extended.Swap(constant.SwapInner),
		),
		html.Select(
			html.ID("scope-select"),
			html.Name(constant.ScopeField),
			extended.Get(constant.EchoPath),
			extended.Trigger(constant.TriggerChange),
			extended.Target(selector(constant.SelectionMark)),
			extended.Swap(constant.SwapInner),
			html.Option(html.Value("first"), gomponents.Text("first")),
			html.Option(html.Value("second"), gomponents.Text("second")),
		),
		html.Div(html.ID(constant.SelectionMark)),
		html.Input(
			html.ID("term-input"),
			html.Type("text"),
			html.Name(constant.TermField),
			extended.Get(constant.EchoPath),
			extended.Trigger(constant.TriggerType),
			extended.Target(selector(constant.SearchMark)),
			extended.Swap(constant.SwapInner),
			extended.Include("#scope-select"),
			extended.Value(`{"origin":"typed"}`),
		),
		html.Div(html.ID(constant.SearchMark)),
	)
}
