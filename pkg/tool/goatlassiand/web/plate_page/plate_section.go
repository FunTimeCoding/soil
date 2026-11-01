package plate_page

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func PlateSection(issues []*issue.Issue) gomponents.Node {
	if len(issues) == 0 {
		return html.P(
			html.Class(webConstant.AbsentClass),
			gomponents.Text(constant.PlateClean),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(constant.PlateTitle)),
			issueTable(issues),
		},
	)
}
