package coverage_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
)

func Region(servers []*coverage.Server) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(servers))

	for _, e := range servers {
		rows = append(rows, row(e))
	}

	return table.New(
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
		Node()
}
