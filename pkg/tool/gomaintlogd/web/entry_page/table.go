package entry_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
)

func Table(entries []entry.Entry) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(entries))

	for _, e := range entries {
		rows = append(rows, opening(&e).Node())
	}

	return table.New(
		"Timestamp",
		"Action",
		"User",
		"System",
		"Service",
		"Description",
	).
		Rows(rows).
		Empty("No entries found.").
		Node()
}
