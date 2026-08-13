package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func impressionWindowLinks(active int) gomponents.Node {
	var nodes []gomponents.Node

	for _, d := range []int{2, 7, 30} {
		if len(nodes) > 0 {
			nodes = append(nodes, gomponents.Text(" · "))
		}

		label := gomponents.Textf("%dd", d)

		if d == active {
			nodes = append(nodes, html.Strong(label))
		} else {
			nodes = append(
				nodes,
				html.A(gomponents.Attr("href", impressionLink(d)), label),
			)
		}
	}

	return gomponents.Group(nodes)
}
