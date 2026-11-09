package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"net/http"
	"slices"
	"strings"
	"testing"
)

func TestEveryOperationClassifiesMissingResource(t *testing.T) {
	c := refusingGithub(t, http.StatusNotFound, nil, `{"message":"Not Found"}`)
	operations := map[string]func() error{
		"ActionRepository": func() error {
			_, e := c.ActionRepository()

			return e
		},
		"DeletePackageVersion": func() error {
			return c.DeletePackageVersion("alfa", 1)
		},
		"DeleteRun": func() error {
			return c.DeleteRun("alfa", "bravo", 1)
		},
		"DeleteTag": func() error {
			return c.DeleteTag("alfa", "bravo", "charlie")
		},
		"FailedRuns": func() error {
			_, e := c.FailedRuns(false)

			return e
		},
		"Issues": func() error {
			_, e := c.Issues()

			return e
		},
		"Jobs": func() error {
			_, e := c.Jobs("alfa", "bravo", 1)

			return e
		},
		"LatestRuns": func() error {
			_, e := c.LatestRuns("alfa", "bravo")

			return e
		},
		"Notifications": func() error {
			_, e := c.Notifications()

			return e
		},
		"OpenUserIssues": func() error {
			_, e := c.OpenUserIssues()

			return e
		},
		"Organizations": func() error {
			_, e := c.Organizations("alfa")

			return e
		},
		"PackageVersions": func() error {
			_, e := c.PackageVersions("alfa")

			return e
		},
		"Packages": func() error {
			_, e := c.Packages("alfa")

			return e
		},
		"ProjectIssues": func() error {
			_, e := c.ProjectIssues("alfa", "bravo")

			return e
		},
		"ProjectRuns": func() error {
			_, e := c.ProjectRuns("alfa", "bravo")

			return e
		},
		"PullRequests": func() error {
			_, e := c.PullRequests("alfa", "bravo")

			return e
		},
		"Reachable": func() error {
			return c.Reachable(t.Context())
		},
		"Releases": func() error {
			_, e := c.Releases("alfa", "bravo")

			return e
		},
		"Repositories": func() error {
			_, e := c.Repositories("alfa")

			return e
		},
		"Runs": func() error {
			_, e := c.Runs(false, false)

			return e
		},
		"SearchCode": func() error {
			_, e := c.SearchCode("alfa")

			return e
		},
		"SearchIssue": func() error {
			_, e := c.SearchIssue("alfa")

			return e
		},
		"SearchRepository": func() error {
			_, e := c.SearchRepository("alfa")

			return e
		},
		"Tags": func() error {
			_, e := c.Tags("alfa", "bravo")

			return e
		},
		"User": func() error {
			_, e := c.User()

			return e
		},
		"Workflows": func() error {
			_, e := c.Workflows("alfa", "bravo")

			return e
		},
	}
	var missed []string

	for name, operation := range operations {
		if !not_found.Is(operation()) {
			missed = append(missed, name)
		}
	}

	slices.Sort(missed)
	assert.String(t, "", strings.Join(missed, ", "))
}

func TestBranchRequestDuringOutageIsUnavailable(t *testing.T) {
	c := refusingGithub(
		t,
		http.StatusServiceUnavailable,
		nil,
		`{"message":"Service Unavailable"}`,
	)
	_, e := c.BranchRequest("alfa", "bravo", "charlie")
	assert.True(t, unavailable.Is(e))
	_, _, f := c.FindBranchRequest("alfa", "bravo", "charlie")
	assert.True(t, unavailable.Is(f))
}
