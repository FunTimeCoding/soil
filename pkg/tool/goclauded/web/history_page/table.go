package history_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/timeline"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
)

func Table(
	entries []*timeline.Entry,
	editable bool,
) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(entries))

	for _, e := range entries {
		rows = append(rows, Row(e, editable))
	}

	columns := []string{"Time", "Event"}

	if editable {
		columns = append(columns, "")
	}

	return table.New(columns...).
		Rows(rows).
		Empty("No events recorded.").
		Node()
}
