package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"testing"
)

func TestChannelDrainLeavesDeferredEntries(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Send(r2.Callsign, r1.Callsign, "deferred message")
	drained, e := s.Service.DrainImmediateQueue("session-1", r1.Callsign)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, drained)
	remaining := s.Check("session-1")
	assert.Count(
		t,
		1,
		fixture.EntriesByKind(remaining.Entries, constant.QueueMessage),
	)
}

func TestChannelDrainTakesImmediateEntries(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.SendImmediate(r2.Callsign, r1.Callsign, "immediate message")
	drained, e := s.Service.DrainImmediateQueue("session-1", r1.Callsign)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, drained)
	assert.String(t, "message", drained[0].Kind)
	remaining := s.Check("session-1")
	assert.Count(
		t,
		0,
		fixture.EntriesByKind(remaining.Entries, constant.QueueMessage),
	)
}

func TestBroadcastIsNeverImmediate(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.SendImmediate(r2.Callsign, "", "broadcast")
	drained, e := s.Service.DrainImmediateQueue("session-1", r1.Callsign)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, drained)
	remaining := s.Check("session-1")
	assert.Count(
		t,
		1,
		fixture.EntriesByKind(remaining.Entries, constant.QueueMessage),
	)
}

func TestRosterActivityIsNeverImmediate(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Announce("session-2", r2.Callsign, "some work", "")
	drained, e := s.Service.DrainImmediateQueue("session-1", r1.Callsign)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, drained)
	remaining := s.Check("session-1")
	assert.Count(
		t,
		1,
		fixture.EntriesByKind(remaining.Entries, constant.QueueSessionAnnounce),
	)
}
