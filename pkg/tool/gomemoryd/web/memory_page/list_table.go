package memory_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_list"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
)

func listTable(l *memory_list.List) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(l.Memories))

	for _, m := range l.Memories {
		rows = append(
			rows,
			opening(m.Identifier, cells(m.Name, m.Description, m.Tags)).Node(),
		)
	}

	return table.New("Name", "Description", "Tags").
		Rows(rows).
		Empty("No memories.").
		Node()
}
