package plate_page

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func IssuesSection(
	title string,
	issues []*issue.Issue,
) gomponents.Node {
	if len(issues) == 0 {
		return gomponents.Group(nil)
	}

	return gomponents.Group(
		[]gomponents.Node{html.H3(gomponents.Text(title)), issueTable(issues)},
	)
}
