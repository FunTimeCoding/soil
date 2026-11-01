package status_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/finding"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(findings []*finding.Finding) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(findings))

	for _, i := range findings {
		subject := gomponents.Text(i.Subject)

		if len(i.Subject) == 36 {
			subject = html.A(
				html.Href(session_page.Path(i.Subject)),
				html.Small(gomponents.Text(i.Subject)),
			)
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(html.Code(gomponents.Text(i.Kind))),
				html.Td(subject),
				html.Td(gomponents.Text(i.Detail)),
			),
		)
	}

	return table.New("Kind", "Subject", "Detail").
		Rows(rows).
		Empty("Nothing inconsistent.").
		Node()
}
