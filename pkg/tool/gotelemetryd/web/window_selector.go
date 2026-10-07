package web

import (
	"fmt"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/selector_option"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func windowSelector(
	currentWindow string,
	currentGroup string,
) gomponents.Node {
	windows := []*selector_option.Option{
		selector_option.New("1h", "1 hour"),
		selector_option.New("24h", "24 hours"),
		selector_option.New("168h", "7 days"),
		selector_option.New("720h", "30 days"),
	}
	groups := []*selector_option.Option{
		selector_option.New(constant.Tool, "by tool"),
		selector_option.New(constant.Surface, "by tool + surface"),
		selector_option.New(constant.Kind, "by tool + kind"),
	}
	var windowLinks []gomponents.Node

	for _, w := range windows {
		link := fmt.Sprintf("/?window=%s&group=%s", w.Value, currentGroup)

		if w.Value == currentWindow {
			windowLinks = append(
				windowLinks,
				html.Strong(gomponents.Text(w.Label)),
			)
		} else {
			windowLinks = append(
				windowLinks,
				html.A(gomponents.Attr("href", link), gomponents.Text(w.Label)),
			)
		}

		windowLinks = append(
			windowLinks,
			gomponents.Text(stringsConstant.SpacedMiddot),
		)
	}

	var groupLinks []gomponents.Node

	for _, g := range groups {
		link := fmt.Sprintf("/?window=%s&group=%s", currentWindow, g.Value)

		if g.Value == currentGroup {
			groupLinks = append(
				groupLinks,
				html.Strong(gomponents.Text(g.Label)),
			)
		} else {
			groupLinks = append(
				groupLinks,
				html.A(gomponents.Attr("href", link), gomponents.Text(g.Label)),
			)
		}

		groupLinks = append(
			groupLinks,
			gomponents.Text(stringsConstant.SpacedMiddot),
		)
	}

	return html.P(
		gomponents.Group(windowLinks[:len(windowLinks)-1]),
		html.Br(),
		gomponents.Group(groupLinks[:len(groupLinks)-1]),
	)
}
