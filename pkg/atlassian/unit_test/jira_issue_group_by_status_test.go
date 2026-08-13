package unit_test

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"testing"
)

func TestGroupByStatus(t *testing.T) {
	o := issue.FixtureOption()
	r1 := issue.Raw("TEST-1")
	r1.Fields.Status = &jira.Status{Name: constant.JiraToDo}
	r2 := issue.Raw("TEST-2")
	r2.Fields.Status = &jira.Status{Name: constant.JiraClosed}
	actual := issue.GroupByStatus(
		[]*issue.Issue{issue.New(r1, o), issue.New(r2, o)},
	)
	assert.Count(t, 2, actual)
	assert.Count(t, 1, actual[constant.JiraToDo])
	assert.String(t, "TEST-1", actual[constant.JiraToDo][0].Key)
	assert.Count(t, 1, actual[constant.JiraClosed])
	assert.String(t, "TEST-2", actual[constant.JiraClosed][0].Key)
}
