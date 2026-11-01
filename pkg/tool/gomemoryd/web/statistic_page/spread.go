package statistic_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/token_summary"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
)

func spread(summary *token_summary.Summary) gomponents.Node {
	return table.New(
		"",
		"Median",
		"Ninetieth",
		"Ninety-ninth",
		"Maximum",
		"Total",
	).
		Numeric(1, 2, 3, 4, 5).
		Rows(
			[]gomponents.Node{
				spreadRow("Block", summary.Block),
				spreadRow("Description", summary.Description),
			},
		).
		Node()
}
