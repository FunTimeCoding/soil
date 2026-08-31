package worker

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func sameIssue(a *issue.Issue, b *issue.Issue) bool {
	return a.Key == b.Key &&
		a.Status == b.Status &&
		a.ChangeTime().Equal(b.ChangeTime())
}
