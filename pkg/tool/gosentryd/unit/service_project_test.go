package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/service_tester"
	"testing"
)

func TestSearchIssuesResolvesTheProjectSlug(t *testing.T) {
	c := service_tester.NewSource(t)
	c.AddProject("41", "backend")
	c.AddProject("42", "frontend")
	c.AddIssue("100", "41", constant.UnresolvedStatus)
	c.AddIssue("200", "42", constant.UnresolvedStatus)
	issues, e := service_tester.Service(t, c).SearchIssues("", "backend", 0, "")
	assert.FatalOnError(t, e)
	assert.String(t, "41", c.IssueSearchProject())
	assert.Count(t, 1, issues)
	assert.String(t, "100", issues[0].Identifier)
}

func TestSearchIssuesWithoutAProjectSearchesEveryProject(t *testing.T) {
	c := service_tester.NewSource(t)
	c.AddProject("41", "backend")
	c.AddIssue("100", "41", constant.UnresolvedStatus)
	c.AddIssue("200", "42", constant.UnresolvedStatus)
	issues, e := service_tester.Service(t, c).SearchIssues("", "", 0, "")
	assert.FatalOnError(t, e)
	assert.String(t, "", c.IssueSearchProject())
	assert.Count(t, 2, issues)
}

func TestSearchEventsResolvesTheProjectSlug(t *testing.T) {
	c := service_tester.NewSource(t)
	c.AddProject("42", "frontend")
	_, e := service_tester.Service(t, c).SearchEvents(
		"level:error",
		"frontend",
		0,
		"",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "42", c.EventSearchProject())
}

func TestSearchRefusesAnUnknownProject(t *testing.T) {
	c := service_tester.NewSource(t)
	c.AddProject("41", "backend")
	_, e := service_tester.Service(t, c).SearchEvents("", "missing", 0, "")
	assert.True(t, not_found.Is(e))
	assert.String(t, "project not found: missing", e.Error())
	_, f := service_tester.Service(t, c).SearchIssues("", "missing", 0, "")
	assert.True(t, not_found.Is(f))
}
