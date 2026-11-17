package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/service_tester"
	"testing"
)

func TestSearchIssuesResolvesProjectSlug(t *testing.T) {
	s := service_tester.New(t)
	s.Source.AddProject("41", "backend")
	s.Source.AddProject("42", "frontend")
	s.Source.AddIssue("100", "41", constant.UnresolvedStatus)
	s.Source.AddIssue("200", "42", constant.UnresolvedStatus)
	issues, e := s.Service().SearchIssues("", "backend", 0, "")
	assert.FatalOnError(t, e)
	assert.String(t, "41", s.Source.IssueSearchProject())
	assert.Count(t, 1, issues)
	assert.String(t, "100", issues[0].Identifier)
}

func TestSearchIssuesWithoutProjectSearchesEveryProject(t *testing.T) {
	s := service_tester.New(t)
	s.Source.AddProject("41", "backend")
	s.Source.AddIssue("100", "41", constant.UnresolvedStatus)
	s.Source.AddIssue("200", "42", constant.UnresolvedStatus)
	issues, e := s.Service().SearchIssues("", "", 0, "")
	assert.FatalOnError(t, e)
	assert.String(t, "", s.Source.IssueSearchProject())
	assert.Count(t, 2, issues)
}

func TestSearchEventsResolvesProjectSlug(t *testing.T) {
	s := service_tester.New(t)
	s.Source.AddProject("42", "frontend")
	_, e := s.Service().SearchEvents("level:error", "frontend", 0, "")
	assert.FatalOnError(t, e)
	assert.String(t, "42", s.Source.EventSearchProject())
}

func TestSearchRefusesUnknownProject(t *testing.T) {
	s := service_tester.New(t)
	s.Source.AddProject("41", "backend")
	_, e := s.Service().SearchEvents("", "missing", 0, "")
	assert.True(t, not_found.Is(e))
	assert.String(t, "project not found: missing", e.Error())
	_, f := s.Service().SearchIssues("", "missing", 0, "")
	assert.True(t, not_found.Is(f))
}
