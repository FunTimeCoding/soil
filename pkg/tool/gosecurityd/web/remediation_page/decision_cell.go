package remediation_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func decisionCell(i *remediation_item.Item) gomponents.Node {
	switch {
	case i.Decision == nil:
		return html.Td(html.Small(gomponents.Text("none")))
	case i.Decision.Mitigated:
		return html.Td(
			html.Strong(gomponents.Text("Mitigated: ")),
			gomponents.Text(i.Decision.Mitigation),
		)
	}

	return html.Td(
		html.Strong(gomponents.Text("Flagged: ")),
		gomponents.Text(i.Decision.Note),
	)
}
