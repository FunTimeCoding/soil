package system_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/package_group"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Kernel(groups []*package_group.Group) gomponents.Node {
	vulnerabilities := 0
	rows := make([]gomponents.Node, 0, len(groups))

	for _, g := range groups {
		vulnerabilities += len(g.Vulnerabilities)
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(g.Kind)),
				html.Td(ScoreBadge(g.TopScore)),
				html.Td(gomponents.Text(g.Installed)),
				html.Td(gomponents.Text(g.Fixed)),
				html.Td(gomponents.Text(g.Candidate)),
				html.Td(
					html.Details(
						html.Summary(
							gomponents.Text(
								fmt.Sprintf(
									"%d vulnerabilities",
									len(g.Vulnerabilities),
								),
							),
						),
						html.Small(
							gomponents.Text(join.CommaSpace(g.Vulnerabilities)),
						),
					),
				),
			),
		)
	}

	return html.Details(
		html.Summary(
			html.Strong(
				gomponents.Text(
					fmt.Sprintf(
						"Kernel: %d vulnerabilities on the running kernel",
						vulnerabilities,
					),
				),
			),
		),
		table.New(
			"Fix",
			"Top score",
			"Running",
			"Fixed up to",
			"Candidate",
			constant.VulnerabilitiesHeader,
		).
			Rows(rows).
			Empty("None.").
			Node(),
	)
}
