package permission_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/permission"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
)

func Region(r *permission.Report) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(r.Rows))

	for _, w := range r.Rows {
		rows = append(rows, row(w))
	}

	return table.New(
		"Blocker",
		"Prompts",
		constant.SessionsTitle,
		"Reason",
		"Last",
	).
		Numeric(1, 2).
		Rows(rows).
		Empty("No permission prompts in this window.").
		Node()
}
