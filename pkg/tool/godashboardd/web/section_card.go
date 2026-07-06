package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func sectionCard(
	v *board.Section,
	values map[string][]string,
) gomponents.Node {
	var nodes []gomponents.Node

	if v.Name != "" {
		nodes = append(nodes, html.H4(gomponents.Text(v.Name)))
	}

	for _, e := range v.Entries {
		nodes = append(nodes, entryItem(e, values))
	}

	return html.Article(
		html.Class("board-section"),
		gomponents.Group(nodes),
	)
}
