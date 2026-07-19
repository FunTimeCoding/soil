package unit_test

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"testing"
)

func TestFilterType(t *testing.T) {
	o := issue.FixtureOption()
	r1 := issue.Raw("TEST-1")
	r1.Fields.Type = jira.IssueType{Name: issue.EpicType}
	r2 := issue.Raw("TEST-2")
	r2.Fields.Type = jira.IssueType{Name: issue.TaskType}
	actual := issue.FilterType(
		[]*issue.Issue{issue.New(r1, o), issue.New(r2, o)},
		issue.TaskType,
	)
	assert.Count(t, 1, actual)
	assert.String(t, "TEST-1", actual[0].Key)
}
