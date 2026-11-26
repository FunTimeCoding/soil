package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"strings"
	"testing"
)

func TestReplyToDiscussionLandsInThread(t *testing.T) {
	s := base.New(t)
	s.Alfa().SeedProject(7, "group", "app")
	result := s.Tools().MustCallTool(
		constant.ReplyToDiscussion,
		map[string]any{
			"project":       "group/app",
			"merge_request": 3,
			"discussion":    gitlabConstant.FixtureDiscussion,
			"body":          "see the runbook",
		},
	)
	assert.StringContains(t, `"Body": "see the runbook"`, result)
	assert.Strings(
		t,
		[]string{"see the runbook"},
		s.Alfa().Replies(gitlabConstant.FixtureDiscussion),
	)
	assert.False(t, s.Alfa().Resolved(gitlabConstant.FixtureDiscussion))
}

func TestReplyToDiscussionResolves(t *testing.T) {
	s := base.New(t)
	s.Alfa().SeedProject(7, "group", "app")
	s.Tools().MustCallTool(
		constant.ReplyToDiscussion,
		map[string]any{
			"project":       "group/app",
			"merge_request": 3,
			"discussion":    gitlabConstant.FixtureDiscussion,
			"body":          "fixed",
			"resolve":       true,
		},
	)
	assert.True(t, s.Alfa().Resolved(gitlabConstant.FixtureDiscussion))
}

func TestReplyToDiscussionRejectsMalformedIdentifier(t *testing.T) {
	s := base.New(t)
	s.Alfa().SeedProject(7, "group", "app")
	assert.StringContains(
		t,
		"discussion must be 40 lowercase hex characters",
		s.Tools().MustCallToolError(
			constant.ReplyToDiscussion,
			map[string]any{
				"project":       "group/app",
				"merge_request": 3,
				"discussion":    "../notes",
				"body":          "x",
			},
		),
	)
	assert.Count(t, 0, s.Alfa().Replies("../notes"))
}

func TestReplyToDiscussionCallsignLeftOutOfOwnWatch(t *testing.T) {
	s := base.New(t, service_tester.WithWatches())
	s.Alfa().SeedProject(7, "group", "app")
	v := s.SeedWatched()
	c := s.Tools()

	for _, callsign := range []string{"kilo", "lima"} {
		c.MustCallTool(
			constant.SubscribeMergeRequest,
			map[string]any{
				"project":       "group/app",
				"merge_request": 3,
				"callsign":      callsign,
			},
		)
	}

	c.MustCallTool(
		constant.ReplyToDiscussion,
		map[string]any{
			"project":       "group/app",
			"merge_request": 3,
			"discussion":    gitlabConstant.FixtureDiscussion,
			"body":          "answered in thread",
			"callsign":      "kilo",
		},
	)
	v.UserNotesCount++
	assert.Count(t, 0, s.Daemon().Service().PollWatches(t.Context()))
	assert.Count(t, 0, s.Bodies("kilo"))
	lima := s.Bodies("lima")
	assert.Count(t, 1, lima)
	assert.True(t, strings.HasSuffix(lima[0], " · kilo: answered in thread"))
}
