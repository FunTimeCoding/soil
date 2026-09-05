package unit

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"testing"
)

func TestFilterType(t *testing.T) {
	o := issue.FixtureOption()
	r1 := issue.Raw("TEST-1")
	r1.Fields.Type = jira.IssueType{Name: constant.JiraEpicType}
	r2 := issue.Raw("TEST-2")
	r2.Fields.Type = jira.IssueType{Name: constant.JiraTaskType}
	actual := issue.FilterType(
		[]*issue.Issue{issue.New(r1, o), issue.New(r2, o)},
		constant.JiraTaskType,
	)
	assert.Count(t, 1, actual)
	assert.String(t, "TEST-1", actual[0].Key)
}
