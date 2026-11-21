package flag_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
)

func Table(
	rows []*overview_result.Row,
	hits map[string]*affected_count.Count,
) gomponents.Node {
	nodes := make([]gomponents.Node, 0, len(rows))

	for _, r := range rows {
		nodes = append(nodes, opening(r, hits[r.Flag.Subject]).Node())
	}

	return table.New("Subject", "Severity", "Hits", "Flagged").
		Rows(nodes).
		Empty("No flags here.").
		Node()
}
