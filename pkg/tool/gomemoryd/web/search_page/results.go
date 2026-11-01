package search_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Results(
	query string,
	results []record.SearchResult,
) gomponents.Node {
	if query == "" {
		return gomponents.Group(nil)
	}

	rows := make([]gomponents.Node, 0, len(results))

	for _, m := range results {
		rows = append(
			rows,
			html.Tr(
				html.Td(memory_page.Link(m.Identifier, m.Name)),
				html.Td(html.Small(gomponents.Text(m.Description))),
				html.Td(gomponents.Text(m.Type)),
				html.Td(memory_page.TagChips(m.Tags)),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			gomponents.If(
				len(results) > 0,
				html.P(gomponents.Text(resultCount(len(results)))),
			),
			table.New("Name", "Description", "Type", "Tags").
				Rows(rows).
				Empty("No results.").
				Node(),
		},
	)
}
