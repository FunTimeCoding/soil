package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/issue"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/google/go-github/v90/github"
	"testing"
)

func TestIssue(t *testing.T) {
	i := issue.New(
		&github.Issue{
			RepositoryURL: locator.New("api.github.com").Path(
				"/repos/funtimecoding/soil",
			).Pointer(),
			Title:   new(constant.UpperAlfa),
			HTMLURL: new(constant.UpperBravo),
		},
	)
	i.Raw = nil
	assert.Any(
		t,
		&issue.Issue{
			Repository: "funtimecoding/soil",
			Title:      "Alfa",
			Link:       "Bravo",
		},
		i,
	)
}
