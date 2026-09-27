package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"testing"
)

func TestHistoryCountEmpty(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "bind identity")
	result := a.MustCallTool(constant.HistoryCount, map[string]any{})
	assert.StringContains(t, "1 events", result)
}

func TestHistoryCountMultiple(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "one")
	a.Announce(a.Name(), "two")
	a.MustCallTool(
		constant.Update,
		map[string]any{constant.Message: "completed", constant.Topic: "three"},
	)
	result := a.MustCallTool(constant.HistoryCount, map[string]any{})
	assert.StringContains(t, "3 events", result)
}

func TestHistoryFormatsIDs(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "test topic")
	history := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "announced: test topic", history)
}

func TestHistoryLimit(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "topic-alfa")
	a.Announce(a.Name(), "topic-bravo")
	a.Announce(a.Name(), "topic-charlie")
	history := a.MustCallTool(
		constant.History,
		map[string]any{constant.Limit: float64(1)},
	)
	assert.StringContains(t, "topic-charlie", history)
	assert.StringNotContains(t, "topic-alfa", history)
}

func TestHistorySkip(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "topic-alfa")
	a.Announce(a.Name(), "topic-bravo")
	a.Announce(a.Name(), "topic-charlie")
	history := a.MustCallTool(
		constant.History,
		map[string]any{constant.Limit: float64(1), constant.Offset: float64(1)},
	)
	assert.StringContains(t, "topic-bravo", history)
	assert.StringNotContains(t, "topic-charlie", history)
}

func TestHistoryAllKinds(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "milestone",
		},
	)
	a.MustCallTool(
		constant.Complete,
		map[string]any{constant.Message: "finished"},
	)
	history := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "announced: working", history)
	assert.StringContains(t, "updated milestone: completed", history)
	assert.StringContains(t, "completed milestone: finished", history)
}

func TestLabelReservedKeyBlocked(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "labeling")

	for _, key := range constant.ReservedLabelKeys {
		result := a.MustCallToolError(
			constant.Label,
			map[string]any{
				constant.Key:   key,
				constant.Value: "should be blocked",
			},
		)
		assert.StringContains(t, "reserved key", result)
		assert.StringContains(t, "edit_session", result)
	}
}

func TestLabelCustomKey(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "labeling")
	a.MustCallTool(
		constant.Label,
		map[string]any{constant.Key: "stage", constant.Value: "canary"},
	)
	result := a.MustCallTool(
		constant.Label,
		map[string]any{constant.Key: "stage"},
	)
	assert.StringContains(t, "stage=canary", result)
}

func TestEditSessionAliasSelf(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{constant.Alias: "my-project"},
	)
	roster := a.MustCallTool(constant.Roster, map[string]any{})
	assert.StringContains(t, "my-project", roster)
}

func TestEditSessionAliasOther(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	b := s.NewSession(t)
	a.Announce(a.Name(), "editor")
	b.Announce(b.Name(), constant.FixtureTarget)
	a.MustCallTool(
		constant.EditSession,
		map[string]any{constant.Alias: "the-target", constant.Target: b.Name()},
	)
	roster := a.MustCallTool(constant.Roster, map[string]any{})
	assert.StringContains(t, "the-target", roster)
}

func TestEditSessionDescription(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{constant.Description: "Fixed the auth bug"},
	)
	assert.String(
		t,
		"Fixed the auth bug",
		s.Store.GetSession(a.UUID).Description,
	)
}

func TestEditSessionBoth(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{
			constant.Alias:       "my-project",
			constant.Description: "Refactored the CLI",
		},
	)
	e := s.Store.GetSession(a.UUID)
	assert.String(t, "my-project", e.AliasValue())
	assert.String(t, "Refactored the CLI", e.Description)
}

func TestSessionStatusWithTopic(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "building search index")
	result := a.MustCallTool(constant.SessionStatus, map[string]any{})
	assert.StringContains(t, a.Name(), result)
	assert.StringContains(t, "building search index", result)
}

func TestSessionStatusAfterComplete(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "some work")
	a.MustCallTool(constant.Complete, map[string]any{constant.Message: "done"})
	result := a.MustCallTool(constant.SessionStatus, map[string]any{})
	assert.StringContains(t, "(none)", result)
}

func TestSessionStatusBeforeAnnounce(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	result := a.MustCallToolError(constant.SessionStatus, map[string]any{})
	assert.StringContains(t, "announce first", result)
}

func TestTimelinePagination(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)

	for i := 1; i <= 12; i++ {
		a.Announce(a.Name(), fmt.Sprintf("topic-%02d", i))
	}

	x := context.Background()
	first, e := a.RestClient.GetTimelineWithResponse(
		x,
		&client.GetTimelineParams{Limit: new(5), Offset: new(0)},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 5, *first.JSON200)
	assert.StringContains(t, "topic-12", (*first.JSON200)[0].Subject)
	assert.StringContains(t, "topic-08", (*first.JSON200)[4].Subject)
	second, f := a.RestClient.GetTimelineWithResponse(
		x,
		&client.GetTimelineParams{Limit: new(5), Offset: new(5)},
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 5, *second.JSON200)
	assert.StringContains(t, "topic-07", (*second.JSON200)[0].Subject)
	assert.StringContains(t, "topic-03", (*second.JSON200)[4].Subject)
	third, g := a.RestClient.GetTimelineWithResponse(
		x,
		&client.GetTimelineParams{Limit: new(5), Offset: new(10)},
	)
	assert.FatalOnError(t, g)
	assert.Count(t, 2, *third.JSON200)
	assert.StringContains(t, "topic-02", (*third.JSON200)[0].Subject)
	assert.StringContains(t, "topic-01", (*third.JSON200)[1].Subject)
}
