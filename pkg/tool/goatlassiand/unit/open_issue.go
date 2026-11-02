package unit

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func openIssue(key string) *issue.Issue {
	r := issue.Raw(key)
	r.Fields.Status = &jira.Status{
		Name:           constant.JiraToDo,
		StatusCategory: jira.StatusCategory{Key: "new"},
	}

	return issue.New(r, issue.FixtureOption())
}
