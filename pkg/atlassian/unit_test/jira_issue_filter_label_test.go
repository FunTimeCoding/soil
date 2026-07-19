package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestFilterLabel(t *testing.T) {
	o := issue.FixtureOption()
	r1 := issue.Raw("TEST-1")
	r1.Fields.Labels = []string{upper.Alfa}
	r2 := issue.Raw("TEST-2")
	r2.Fields.Labels = []string{upper.Bravo}
	actual := issue.FilterLabel(
		[]*issue.Issue{issue.New(r1, o), issue.New(r2, o)},
		upper.Bravo,
	)
	assert.Count(t, 1, actual)
	assert.String(t, "TEST-1", actual[0].Key)
}
