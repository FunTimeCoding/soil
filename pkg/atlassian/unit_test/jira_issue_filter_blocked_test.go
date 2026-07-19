package unit_test

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"testing"
)

func TestFilterBlocked(t *testing.T) {
	o := issue.FixtureOption()
	r2 := issue.Raw("TEST-2")
	blockBy(r2, "TEST-99", constant.InProgress)
	r3 := issue.Raw("TEST-3")
	blockBy(r3, "TEST-100", constant.Closed)
	actual := issue.FilterBlocked(
		[]*issue.Issue{
			issue.New(issue.Raw("TEST-1"), o),
			issue.New(r2, o),
			issue.New(r3, o),
		},
	)
	assert.Count(t, 2, actual)
	assert.String(t, "TEST-1", actual[0].Key)
	assert.String(t, "TEST-3", actual[1].Key)
}

func blockBy(
	i *jira.Issue,
	key string,
	status string,
) {
	i.Fields.IssueLinks = append(
		i.Fields.IssueLinks,
		&jira.IssueLink{
			Type: jira.IssueLinkType{Inward: issue.BlockedBy},
			InwardIssue: &jira.Issue{
				Key:    key,
				Fields: &jira.IssueFields{Status: &jira.Status{Name: status}},
			},
		},
	)
}
