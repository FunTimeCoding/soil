package coverage_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Region(servers []*coverage.Server) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(servers))
	var details []gomponents.Node

	for _, e := range servers {
		name := e.Name

		if !e.Configured {
			name = fmt.Sprintf("%s (unconfigured)", name)
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(name)),
				html.Td(gomponents.Text(ratio(e.UsedRecent, e.Registered))),
				html.Td(gomponents.Text(ratio(e.UsedTotal, e.Registered))),
				numeric(e.CallsRecent),
				numeric(e.CallsTotal),
				lastUsed(e.LastUsed),
			),
		)

		if len(e.Tools) > 0 {
			details = append(
				details,
				html.Details(html.Summary(gomponents.Text(e.Name)), tools(e)),
			)
		}
	}

	return gomponents.Group(
		[]gomponents.Node{
			table.New(
				"Server",
				"Coverage 30d",
				"Coverage all",
				"Calls 30d",
				"Calls all",
				"Last used",
			).
				Numeric(3, 4).
				Rows(rows).
				Empty("No coverage data.").
				Node(),
			gomponents.Group(details),
		},
	)
}
