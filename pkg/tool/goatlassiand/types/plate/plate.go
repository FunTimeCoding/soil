package plate

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

type Plate struct {
	Issues        []*issue.Issue
	WatchedIssues []*issue.Issue
	Newest        []*issue.Issue
	Favorites     []*page.Page
	Watched       []*page.Page
}
