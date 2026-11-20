package flag_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
)

func Table(rows []*overview_result.Row) gomponents.Node {
	nodes := make([]gomponents.Node, 0, len(rows))

	for _, r := range rows {
		nodes = append(nodes, opening(r).Node())
	}

	return table.New("CVE", "Severity", "Debian", "Note", "Flagged").
		Rows(nodes).
		Empty("No flags here.").
		Node()
}
