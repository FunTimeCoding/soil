package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestOnlyInitials(t *testing.T) {
	assert.Any(
		t,
		[]*issue.Issue{{Initials: "ALFA"}},
		issue.OnlyInitials(
			[]*issue.Issue{
				{Initials: stringsConstant.CapitalAlfa},
				{Initials: stringsConstant.CapitalBravo},
			},
			stringsConstant.CapitalAlfa,
		),
	)
}
