package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/atlassian/unit/atlassian_tester"
	"testing"
)

func TestFilterBlocked(t *testing.T) {
	o := issue.FixtureOption()
	r2 := issue.Raw("TEST-2")
	atlassian_tester.BlockBy(r2, "TEST-99", constant.JiraInProgress)
	r3 := issue.Raw("TEST-3")
	atlassian_tester.BlockBy(r3, "TEST-100", constant.JiraClosed)
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
