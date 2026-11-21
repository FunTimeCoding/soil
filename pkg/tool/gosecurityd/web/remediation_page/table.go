package remediation_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/system_page"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(
	items []*remediation_item.Item,
	empty string,
) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(items))

	for _, i := range items {
		vulnerabilities := listCell(i.Vulnerabilities, "vulnerabilities")

		if i.Unknown {
			vulnerabilities = html.Td(html.Small(gomponents.Text("unknown")))
		}

		rows = append(
			rows,
			html.Tr(
				itemCell(i),
				kindCell(i),
				listCell(i.Hosts, "hosts"),
				vulnerabilities,
				html.Td(system_page.ScoreBadge(i.TopScore)),
				decisionCell(i),
			),
		)
	}

	return table.New(
		"Item",
		"Kind",
		"Hosts",
		constant.VulnerabilitiesHeader,
		"Top score",
		"Decision",
	).
		Rows(rows).
		Empty(empty).
		Node()
}
