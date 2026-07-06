package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"testing"
)

func TestFilterBlocked(t *testing.T) {
	o := fixtureOption()
	r2 := Raw("TEST-2")
	blockBy(r2, "TEST-99", constant.InProgress)
	r3 := Raw("TEST-3")
	blockBy(r3, "TEST-100", constant.Closed)
	actual := FilterBlocked(
		[]*Issue{
			New(Raw("TEST-1"), o),
			New(r2, o),
			New(r3, o),
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
			Type: jira.IssueLinkType{Inward: BlockedBy},
			InwardIssue: &jira.Issue{
				Key:    key,
				Fields: &jira.IssueFields{Status: &jira.Status{Name: status}},
			},
		},
	)
}
