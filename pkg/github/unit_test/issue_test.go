package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/issue"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/google/go-github/v89/github"
	"testing"
)

func TestIssue(t *testing.T) {
	i := issue.New(
		&github.Issue{
			RepositoryURL: locator.New(
				"api.github.com",
			).Path("/repos/funtimecoding/soil").Pointer(),
			Title:   new(upper.Alfa),
			HTMLURL: new(upper.Bravo),
		},
	)
	i.Raw = nil
	assert.Any(
		t,
		&issue.Issue{
			Repository: "funtimecoding/soil",
			Title:      upper.Alfa,
			Link:       upper.Bravo,
		},
		i,
	)
}
