package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestChildren(t *testing.T) {
	actual := Children(
		[]*Issue{
			{
				Key: "TEST-1",
				Raw: &jira.Issue{Fields: &jira.IssueFields{}},
			},
			{
				Key: "TEST-2",
				Raw: &jira.Issue{
					Fields: &jira.IssueFields{
						Parent: &jira.Parent{Key: "TEST-1"},
					},
				},
			},
			{
				Key: "TEST-3",
				Raw: &jira.Issue{
					Fields: &jira.IssueFields{
						Parent: &jira.Parent{Key: "TEST-99"},
					},
				},
			},
		},
		"TEST-1",
	)
	assert.Count(t, 1, actual)
	assert.String(t, "TEST-2", actual[0].Key)
}
