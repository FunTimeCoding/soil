package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/strings/capital"
	"testing"
)

func TestOnlyInitials(t *testing.T) {
	assert.Any(
		t,
		[]*issue.Issue{{Initials: "ALFA"}},
		issue.OnlyInitials(
			[]*issue.Issue{
				{Initials: capital.Alfa},
				{Initials: capital.Bravo},
			},
			capital.Alfa,
		),
	)
}
