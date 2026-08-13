package web

import (
	"fmt"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func windowSelector(
	currentWindow string,
	currentGroup string,
) gomponents.Node {
	windows := []SelectorOption{
		{"1h", "1 hour"},
		{"24h", "24 hours"},
		{"168h", "7 days"},
		{"720h", "30 days"},
	}
	groups := []SelectorOption{
		{"tool", "by tool"},
		{"surface", "by tool + surface"},
		{"kind", "by tool + kind"},
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

		windowLinks = append(windowLinks, gomponents.Text(" · "))
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

		groupLinks = append(groupLinks, gomponents.Text(" · "))
	}

	return html.P(
		gomponents.Group(windowLinks[:len(windowLinks)-1]),
		html.Br(),
		gomponents.Group(groupLinks[:len(groupLinks)-1]),
	)
}
