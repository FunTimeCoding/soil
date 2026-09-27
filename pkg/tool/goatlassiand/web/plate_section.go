package web

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func plateSection(issues []*issue.Issue) gomponents.Node {
	if len(issues) == 0 {
		return html.P(gomponents.Text(constant.PlateClean))
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(constant.PlateTitle)),
			issuesTable(issues),
		},
	)
}
