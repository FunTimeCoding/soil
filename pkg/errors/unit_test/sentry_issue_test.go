package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestIssue(t *testing.T) {
	r := response.NewIssue()
	r.Identifier = "1"
	r.Type = constant.UpperBravo
	r.Title = constant.UpperCharlie
	r.Permalink = constant.UpperDelta
	r.Project.Name = constant.UpperAlfa
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
