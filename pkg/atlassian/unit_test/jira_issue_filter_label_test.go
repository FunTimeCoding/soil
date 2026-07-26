package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestFilterLabel(t *testing.T) {
	o := issue.FixtureOption()
	r1 := issue.Raw("TEST-1")
	r1.Fields.Labels = []string{constant.UpperAlfa}
	r2 := issue.Raw("TEST-2")
	r2.Fields.Labels = []string{constant.UpperBravo}
	actual := issue.FilterLabel(
		[]*issue.Issue{issue.New(r1, o), issue.New(r2, o)},
		constant.UpperBravo,
	)
	assert.Count(t, 1, actual)
	assert.String(t, "TEST-1", actual[0].Key)
}
