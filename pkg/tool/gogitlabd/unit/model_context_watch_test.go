package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"strings"
	"testing"
)

func TestSubscribeMergeRequestTool(t *testing.T) {
	s := base.New(t, service_tester.WithWatches())
	s.Alfa().SeedProject(7, "group", "app")
	s.SeedWatched()
	c := s.Tools()
	result := c.MustCallTool(
		constant.SubscribeMergeRequest,
		map[string]any{
			"project":       "group/app",
			"merge_request": 3,
			"callsign":      "kilo",
			"alias":         "app change",
			"delivery":      "queued",
		},
	)
	assert.StringContains(t, `"reference": "group/app!3"`, result)
	assert.StringContains(t, `"alias": "app change"`, result)
	assert.StringContains(t, `"title": "watched change"`, result)
	assert.StringContains(t, `"merge_status": "mergeable"`, result)
	assert.StringContains(t, `"pipeline_status": "running"`, result)
	assert.StringContains(t, `"delivery": "queued"`, result)
	assert.String(
		t,
		"already subscribed",
		c.MustCallTool(
			constant.SubscribeMergeRequest,
			map[string]any{
				"project":       "group/app",
				"merge_request": 3,
				"callsign":      "kilo",
			},
		),
	)
	listed := c.MustCallTool(
		constant.ListSubscriptions,
		map[string]any{"callsign": "kilo"},
	)
	assert.StringContains(t, `"reference": "group/app!3"`, listed)
	assert.StringContains(t, `"instance": "alfa"`, listed)
	assert.String(
		t,
		"unsubscribed",
		c.MustCallTool(
			constant.UnsubscribeMergeRequest,
			map[string]any{
				"project":       "group/app",
				"merge_request": 3,
				"callsign":      "kilo",
			},
		),
	)
	assert.False(
		t,
		strings.Contains(
			c.MustCallTool(
				constant.ListSubscriptions,
				map[string]any{"callsign": "kilo"},
			),
			constant.FixtureReference,
		),
	)
}

func TestSubscribeMergeRequestToolDelivery(t *testing.T) {
	s := base.New(t, service_tester.WithWatches())
	s.Alfa().SeedProject(7, "group", "app")
	s.SeedWatched()
	assert.StringContains(
		t,
		"value must be one of 'immediate', 'queued'",
		s.Tools().MustCallToolError(
			constant.SubscribeMergeRequest,
			map[string]any{
				"project":       "group/app",
				"merge_request": 3,
				"callsign":      "kilo",
				"delivery":      "later",
			},
		),
	)
}

func TestSubscribeMergeRequestToolWithoutGoclauded(t *testing.T) {
	s := base.New(t)
	s.Alfa().SeedProject(7, "group", "app")
	s.SeedWatched()
	assert.StringContains(
		t,
		"merge request watches need goclauded",
		s.Tools().MustCallToolError(
			constant.SubscribeMergeRequest,
			map[string]any{
				"project":       "group/app",
				"merge_request": 3,
				"callsign":      "kilo",
			},
		),
	)
}

func TestCreateMergeRequestNoteToolCallsign(t *testing.T) {
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
		constant.CreateMergeRequestNote,
		map[string]any{
			"project":       "group/app",
			"merge_request": 3,
			"body":          "on it",
			"callsign":      "kilo",
		},
	)
	v.UserNotesCount++
	assert.Count(t, 0, s.Daemon().Service().PollWatches(t.Context()))
	assert.Count(t, 0, s.Bodies("kilo"))
	lima := s.Bodies("lima")
	assert.Count(t, 1, lima)
	assert.True(t, strings.HasSuffix(lima[0], " · kilo: on it"))
}
