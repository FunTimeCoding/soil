package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/event_tester"
	"net/http"
	"testing"
)

func TestEventStreamReplaysFromZero(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "first topic")
	events, disconnect := event_tester.SubscribeStream(
		t,
		s,
		"consumer",
		constant.Announce,
	)
	defer disconnect()
	e := event_tester.Next(t, events)
	assert.String(t, "entry", e.Name)
	assert.StringContains(t, "first topic", e.Payload)
	assert.True(t, e.Identifier != "")
}

func TestEventStreamFiltersKinds(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "announced topic")
	a.MustCallTool(
		constant.Label,
		map[string]any{constant.Key: "colour", constant.Value: "green"},
	)
	events, disconnect := event_tester.SubscribeStream(
		t,
		s,
		"consumer",
		constant.Label,
	)
	defer disconnect()
	e := event_tester.Next(t, events)
	assert.StringContains(t, "green", e.Payload)
	assert.StringNotContains(t, "announced topic", e.Payload)
}

func TestEventStreamResumesFromStoredPosition(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "earlier topic")
	first, disconnect := event_tester.SubscribeStream(
		t,
		s,
		"consumer",
		constant.Announce,
	)
	event_tester.Next(t, first)
	disconnect()
	a.Announce(a.Name(), "later topic")
	second, stop := event_tester.SubscribeStream(
		t,
		s,
		"consumer",
		constant.Announce,
	)
	defer stop()
	e := event_tester.Next(t, second)
	assert.StringContains(t, "later topic", e.Payload)
	assert.StringNotContains(t, "earlier topic", e.Payload)
}

func TestEventStreamRejectsMissingToken(t *testing.T) {
	s := base.New(t)
	assert.Integer(
		t,
		http.StatusUnauthorized,
		event_tester.StatusWithoutToken(t, s),
	)
}
