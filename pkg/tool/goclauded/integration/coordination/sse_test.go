package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/fixture"
	"testing"
)

func TestSSEPushesRosterOnConnect(t *testing.T) {
	s := base.New(t)
	events, disconnect := fixture.Subscribe(t, s)
	defer disconnect()
	e := fixture.NextEvent(t, events)
	assert.String(t, "roster", e.Name)
	assert.True(t, e.Payload != "")
}

func TestSSEPushesActivityOnConnect(t *testing.T) {
	s := base.New(t)
	events, disconnect := fixture.Subscribe(t, s)
	defer disconnect()
	fixture.NextEvent(t, events)
	e := fixture.NextEvent(t, events)
	assert.String(t, "activity", e.Name)
}

func TestSSEPushesAfterMutation(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	events, disconnect := fixture.Subscribe(t, s)
	defer disconnect()
	fixture.NextEvent(t, events)
	fixture.NextEvent(t, events)
	a.Announce(a.Name(), "trigger notification")
	e := fixture.NextEvent(t, events)
	assert.String(t, "roster", e.Name)
	assert.StringContains(t, "trigger notification", e.Payload)
}

func TestSSEMultiLineBody(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "setup")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{constant.Body: "line one\nline two\nline three"},
	)
	events, disconnect := fixture.Subscribe(t, s)
	defer disconnect()
	fixture.NextEvent(t, events)
	activity := fixture.NextEvent(t, events)
	assert.String(t, "activity", activity.Name)
	assert.StringContains(t, "line one", activity.Payload)
	assert.StringContains(t, "line two", activity.Payload)
}
