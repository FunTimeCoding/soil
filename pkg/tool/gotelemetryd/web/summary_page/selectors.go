package summary_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/selector_option"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Selectors(
	window string,
	group string,
) gomponents.Node {
	windows := []*selector_option.Option{
		selector_option.New("1h", "1 hour"),
		selector_option.New(constant.DefaultWindow, "24 hours"),
		selector_option.New("168h", "7 days"),
		selector_option.New("720h", "30 days"),
	}
	groups := []*selector_option.Option{
		selector_option.New(constant.Tool, "by tool"),
		selector_option.New(constant.Surface, "by tool + surface"),
		selector_option.New(constant.Kind, "by tool + kind"),
	}
	var windowChips []gomponents.Node

	for _, w := range windows {
		windowChips = append(
			windowChips,
			choice(w.Label, Path(w.Value, group), w.Value == window),
		)
	}

	var groupChips []gomponents.Node

	for _, g := range groups {
		groupChips = append(
			groupChips,
			choice(g.Label, Path(window, g.Value), g.Value == group),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.Div(
				html.Class(webConstant.ClusterClass),
				gomponents.Group(windowChips),
			),
			html.Div(
				html.Class(webConstant.ClusterClass),
				gomponents.Group(groupChips),
			),
		},
	)
}
