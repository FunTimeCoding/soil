package plate_page

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/web/component/status_dot"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func issueTable(issues []*issue.Issue) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(issues))

	for _, i := range issues {
		rows = append(
			rows,
			html.Tr(
				html.Td(status_dot.New(statusColor(i)).Title(i.Status).Node()),
				html.Td(
					html.A(
						html.Href(i.Link),
						html.Target("_blank"),
						gomponents.Text(i.Key),
					),
				),
				html.Td(
					html.Class(constant.StatusClass),
					gomponents.Text(i.Status),
				),
				html.Td(gomponents.Text(i.Summary)),
				layout.TimeCell(i.ChangeTime()),
			),
		)
	}

	return table.New("", "Key", "Status", "Summary", "Changed").
		Class(constant.IssueTableClass).
		Rows(rows).
		Node()
}
