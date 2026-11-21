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

func Packages(
	title string,
	groups []*package_group.Group,
	open bool,
) gomponents.Node {
	vulnerabilities := 0
	rows := make([]gomponents.Node, 0, len(groups))

	for _, g := range groups {
		vulnerabilities += len(g.Vulnerabilities)
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(g.Source)),
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
		gomponents.If(open && len(groups) > 0, html.Open()),
		html.Summary(
			html.Strong(
				gomponents.Text(
					fmt.Sprintf(
						"%s: %d packages, %d vulnerabilities",
						title,
						len(groups),
						vulnerabilities,
					),
				),
			),
		),
		table.New(
			"Package",
			"Top score",
			"Installed",
			"Fixed in",
			"Candidate",
			constant.VulnerabilitiesHeader,
		).
			Rows(rows).
			Empty("None.").
			Node(),
	)
}
