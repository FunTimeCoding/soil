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
				kernelCell(v),
				criticalCell(v),
				html.Td(gomponents.Text(strconv.Itoa(v.Upgradable))),
				html.Td(gomponents.Text(strconv.Itoa(v.PaidOnly))),
				html.Td(html.Small(gomponents.Text(strconv.Itoa(v.NoFix)))),
				layout.OptionalTimeCell(v.ListsUpdated, "not reported"),
				reported,
			),
		)
	}

	return table.New(
		"System",
		"Release",
		"Coverage",
		"Kernel",
		"9.0 or higher",
		"Package upgrade fixes",
		"Paid only",
		"No fix",
		"Package lists",
		"Reported",
	).
		Numeric(4, 5, 6, 7).
		Rows(nodes).
		Empty("No systems have reported.").
		Node()
}
