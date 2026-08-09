package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func relationTypeFilter(
	current string,
	total int,
	untyped int,
	typeCounts map[string]int,
) gomponents.Node {
	entries := []gomponents.Node{
		html.Small(gomponents.Text("Type: ")),
		relationFilterLink(
			fmt.Sprintf("all (%d)", total),
			"",
			current,
		),
		gomponents.Text(" · "),
		relationFilterLink(
			fmt.Sprintf("untyped (%d)", untyped),
			constant.UntypedFilter,
			current,
		),
	}

	for _, name := range constant.RelationTypes {
		if typeCounts[name] == 0 {
			continue
		}

		entries = append(
			entries,
			gomponents.Text(" · "),
			relationFilterLink(
				fmt.Sprintf("%s (%d)", name, typeCounts[name]),
				name,
				current,
			),
		)
	}

	return html.P(entries...)
}
