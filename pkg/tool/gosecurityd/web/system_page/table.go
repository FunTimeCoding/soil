package system_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strconv"
)

func Table(rows []*system_result.Summary) gomponents.Node {
	nodes := make([]gomponents.Node, 0, len(rows))

	for _, v := range rows {
		reported := layout.TimeCell(v.Updated)

		if v.Stale {
			reported = html.Td(
				badge.New("stale").Color(webConstant.WarnColor).Node(),
			)
		}

		nodes = append(
			nodes,
			html.Tr(
				html.Td(
					html.A(html.Href(link(v.Name)), gomponents.Text(v.Name)),
				),
				html.Td(gomponents.Text(v.Release)),
				html.Td(CoverageBadge(v.Coverage)),
				html.Td(gomponents.Text(strconv.Itoa(v.Upgradable))),
				html.Td(gomponents.Text(strconv.Itoa(v.PaidOnly))),
				html.Td(gomponents.Text(strconv.Itoa(v.NoFix))),
				reported,
			),
		)
	}

	return table.New(
		"System",
		"Release",
		"Coverage",
		"Upgrade fixes",
		"Paid only",
		"No fix",
		"Reported",
	).
		Numeric(3, 4, 5).
		Rows(nodes).
		Empty("No systems have reported.").
		Node()
}
