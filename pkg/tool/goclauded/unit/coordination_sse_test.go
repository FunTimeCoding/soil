package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/event_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"testing"
)

func TestSSEPushesRosterOnConnect(t *testing.T) {
	s := base.New(t)
	v := event_tester.New(t)
	events, disconnect := v.Subscribe(s)
	defer disconnect()
	e := v.Next(events)
	assert.String(t, "roster", e.Name)
	assert.True(t, e.Payload != "")
}

func TestSSEPushesActivityOnConnect(t *testing.T) {
	s := base.New(t)
	v := event_tester.New(t)
	events, disconnect := v.Subscribe(s)
	defer disconnect()
	v.Next(events)
	e := v.Next(events)
	assert.String(t, "activity", e.Name)
}

func TestSSEPushesAfterMutation(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	v := event_tester.New(t)
	events, disconnect := v.Subscribe(s)
	defer disconnect()
	v.NextNamed(events, constant.Activity)
	a.Announce(a.Name(), "trigger notification")
	e := v.NextNamed(events, constant.Roster)
	assert.StringContains(t, "trigger notification", e.Payload)
}

func TestSSEPushesEachRosterCardsState(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "carded topic")
	v := event_tester.New(t)
	events, disconnect := v.Subscribe(s)
	defer disconnect()
	e := v.NextNamed(
		events,
		session_page.Event(a.UUID, constant.RosterStatePart),
	)
	assert.StringContains(t, "carded topic", e.Payload)
}

func TestSSEMultiLineBody(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "setup")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{constant.Body: "line one\nline two\nline three"},
	)
	v := event_tester.New(t)
	events, disconnect := v.Subscribe(s)
	defer disconnect()
	activity := v.NextNamed(events, constant.Activity)
	assert.StringContains(t, "line one", activity.Payload)
	assert.StringContains(t, "line two", activity.Payload)
}
