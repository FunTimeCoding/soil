package floor_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/guest"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/node"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func section(
	n node.Node,
	guests []guest.Guest,
	links map[guest.Guest]string,
) []gomponents.Node {
	var rows []gomponents.Node

	for _, g := range guests {
		if g.Hypervisor != n.Hypervisor || g.Node != n.Name {
			continue
		}

		rows = append(rows, row(g, links[g]))
	}

	return []gomponents.Node{
		html.Div(
			html.Class(constant.NodeHeaderClass),
			html.H3(gomponents.Text(label(n))),
			html.Span(
				html.Class(constant.ReleaseClass),
				gomponents.Text(n.Version),
			),
			gomponents.If(
				n.UpdatesPending > 0,
				badge.New(fmt.Sprintf("%d updates", n.UpdatesPending)).
					Color(webConstant.WarnColor).
					Node(),
			),
		),
		table.New().
			Headerless().
			Class(constant.FloorTableClass).
			Rows(rows).
			Node(),
	}
}
