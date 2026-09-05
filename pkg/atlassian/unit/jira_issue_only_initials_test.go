package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestOnlyInitials(t *testing.T) {
	assert.Any(
		t,
		[]*issue.Issue{{Initials: "ALFA"}},
		issue.OnlyInitials(
			[]*issue.Issue{
				{Initials: constant.CapitalAlfa},
				{Initials: constant.CapitalBravo},
			},
			constant.CapitalAlfa,
		),
	)
}
