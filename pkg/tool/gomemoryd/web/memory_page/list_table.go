package memory_page

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_list"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func listTable(l *memory_list.List) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(l.Memories))

	for _, m := range l.Memories {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						html.Href(Path(m.Identifier)),
						gomponents.Text(m.Name),
					),
				),
				html.Td(html.Small(gomponents.Text(m.Description))),
				html.Td(gomponents.Text(m.Type)),
				html.Td(TagChips(m.Tags)),
				layout.TimeCell(library.Parse(time.RFC3339, m.UpdatedAt)),
			),
		)
	}

	return table.New("Name", "Description", "Type", "Tags", "Updated").
		Rows(rows).
		Empty("No memories.").
		Node()
}
