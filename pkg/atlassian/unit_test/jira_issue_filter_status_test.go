package unit_test

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"testing"
)

func TestFilterStatus(t *testing.T) {
	o := issue.FixtureOption()
	r1 := issue.Raw("TEST-1")
	r1.Fields.Status = &jira.Status{Name: constant.ToDo}
	r2 := issue.Raw("TEST-2")
	r2.Fields.Status = &jira.Status{Name: constant.Closed}
	actual := issue.FilterStatus(
		[]*issue.Issue{issue.New(r1, o), issue.New(r2, o)},
		constant.Closed,
	)
	assert.Count(t, 1, actual)
	assert.String(t, "TEST-1", actual[0].Key)
}
