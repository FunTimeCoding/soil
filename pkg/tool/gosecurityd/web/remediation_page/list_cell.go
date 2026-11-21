package remediation_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func listCell(
	values []string,
	noun string,
) gomponents.Node {
	return html.Td(
		html.Details(
			html.Summary(
				gomponents.Text(fmt.Sprintf("%d %s", len(values), noun)),
			),
			html.Small(gomponents.Text(join.CommaSpace(values))),
		),
	)
}
