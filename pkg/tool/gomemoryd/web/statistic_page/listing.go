package statistic_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/token_summary"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func listing(summary *token_summary.Summary) gomponents.Node {
	var rows []gomponents.Node

	for _, c := range summary.Statistic {
		if c.Hidden {
			continue
		}

		rows = append(
			rows,
			html.Tr(
				numericCell(c.Block),
				numericCell(c.Description),
				html.Td(memory_page.Link(c.Identifier, c.Name)),
				html.Td(gomponents.Text(join.Space(c.Tags...))),
			),
		)
	}

	return table.New("Block", "Description", "Memory", "Tags").
		Numeric(0, 1).
		Rows(rows).
		Empty("No memories.").
		Node()
}
