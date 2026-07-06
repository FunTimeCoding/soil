package issue

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/google/go-github/v88/github"
	"testing"
)

func TestIssue(t *testing.T) {
	i := New(
		&github.Issue{
			RepositoryURL: locator.New(
				"api.github.com",
			).Path("/repos/funtimecoding/go-library").Pointer(),
			Title:   new(upper.Alfa),
			HTMLURL: new(upper.Bravo),
		},
	)
	i.Raw = nil
	assert.Any(
		t,
		&Issue{
			Repository: "funtimecoding/go-library",
			Title:      upper.Alfa,
			Link:       upper.Bravo,
		},
		i,
	)
}
