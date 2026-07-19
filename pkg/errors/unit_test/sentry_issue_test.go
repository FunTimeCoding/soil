package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestIssue(t *testing.T) {
	r := response.NewIssue()
	r.ID = "1"
	r.Type = upper.Bravo
	r.Title = upper.Charlie
	r.Permalink = upper.Delta
	r.Project.Name = upper.Alfa
	actual := issue.New(r)
	actual.Create = nil
	actual.Raw = nil
	assert.Any(
		t,
		&issue.Issue{
			MonitorIdentifier: "sentry-1",
			Project:           "Alfa",
			Type:              "Bravo",
			Title:             "Charlie",
			Link:              "Delta",
		},
		actual,
	)
}
