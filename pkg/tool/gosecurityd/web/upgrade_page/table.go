package upgrade_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_group"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
)

func Table(groups []*upgrade_group.Group) gomponents.Node {
	nodes := make([]gomponents.Node, 0, len(groups))

	for _, g := range groups {
		nodes = append(nodes, opening(g).Node())
	}

	return table.New(
		"Package",
		"Release",
		"Top score",
		"Hosts",
		constant.VulnerabilitiesHeader,
		"Installed",
		"Candidate",
	).
		Numeric(3, 4).
		Rows(nodes).
		Empty("No upgrade fixes anything here.").
		Node()
}
