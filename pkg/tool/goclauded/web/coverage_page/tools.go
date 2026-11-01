package coverage_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func tools(e *coverage.Server) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(e.Tools))

	for _, t := range e.Tools {
		state := badge.New("registered")

		if !t.Registered {
			state = badge.New("observed").Quiet()

			if e.Registered > 0 {
				state = badge.New("retired").Quiet()
			}
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(t.Name)),
				html.Td(state.Node()),
				numeric(t.CallsRecent),
				numeric(t.CallsTotal),
				lastUsed(t.LastUsed),
			),
		)
	}

	return table.New("Tool", "State", "Calls 30d", "Calls all", "Last used").
		Numeric(2, 3).
		Rows(rows).
		Node()
}
