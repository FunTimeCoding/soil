package upgrade_page

import (
	"fmt"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_group"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maps"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
	"slices"
)

func Detail(g *upgrade_group.Group) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(g.Hosts))

	for _, h := range g.Hosts {
		labels := make([]string, 0, len(h.Labels))

		for _, k := range slices.Sorted(maps.Keys(h.Labels)) {
			labels = append(labels, h.Labels[k])
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						html.Href(
							join.Empty(
								constant.SystemPath,
								"?",
								url.Values{
									generativeConstant.ParameterName: {h.Name},
								}.Encode(),
							),
						),
						gomponents.Text(h.Name),
					),
				),
				html.Td(html.Small(gomponents.Text(join.CommaSpace(labels)))),
				html.Td(html.Small(gomponents.Text(h.Installed))),
				html.Td(html.Small(gomponents.Text(h.Candidate))),
				html.Td(html.Code(gomponents.Text(aptLine(h.Binaries)))),
			),
		)
	}

	return opening(g).Detail(
		constant.UpgradeColumns,
		table.New("Host", "Labels", "Installed", "Candidate", "Upgrade").
			Rows(rows).
			Node(),
		gomponents.If(
			g.Kernel,
			html.P(
				html.Strong(
					gomponents.Text(
						"Then reboot into the new kernel; the upgrade alone fixes nothing.",
					),
				),
			),
		),
		html.Details(
			html.Summary(
				gomponents.Text(
					fmt.Sprintf(
						"%d vulnerabilities, fixed up to %s",
						len(g.Vulnerabilities),
						g.Fixed,
					),
				),
			),
			html.Small(gomponents.Text(join.CommaSpace(g.Vulnerabilities))),
		),
	).Response()
}
