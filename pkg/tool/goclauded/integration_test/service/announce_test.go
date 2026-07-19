package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestAnnounce(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Announce("session-1", r.Callsign, "fixing auth", "pkg/auth")
	e := s.Store.GetSession("session-1")
	assert.String(t, "fixing auth", e.Topic)
	assert.String(t, "pkg/auth", e.Files)
	events := s.Store.Events(event_query.New().Kind(constant.Announce).SetLimit(10))
	assert.Count(t, 1, events)
	assert.String(t, "fixing auth", events[0].Metadata[constant.Topic])
}

func TestAnnounceUpdatesTopic(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Announce("session-1", r.Callsign, "first topic", "")
	s.Announce("session-1", r.Callsign, "second topic", "")
	assert.String(
		t,
		"second topic",
		s.Store.GetSession("session-1").Topic,
	)
}
