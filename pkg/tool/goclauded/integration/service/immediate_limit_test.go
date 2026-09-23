package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"testing"
	"time"
)

func TestImmediateIsRateLimitedToTheSlowPath(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")

	for range constant.ImmediateLimit {
		delivered, e := s.Service.Send(r2.Callsign, r1.Callsign, "burst", true)
		assert.FatalOnError(t, e)
		assert.True(t, delivered)
	}

	delivered, e := s.Service.Send(r2.Callsign, r1.Callsign, "over", true)
	assert.FatalOnError(t, e)
	assert.False(t, delivered)
	drained, f := s.Service.DrainImmediateQueue("session-1", r1.Callsign)
	assert.FatalOnError(t, f)
	assert.Count(t, 6, drained)
	remaining := s.Check("session-1")
	assert.Count(
		t,
		1,
		fixture.EntriesByKind(remaining.Entries, constant.QueueMessage),
	)
}

func TestImmediateRecoversAfterTheWindow(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")

	for range constant.ImmediateLimit {
		_, e := s.Service.Send(r2.Callsign, r1.Callsign, "burst", true)
		assert.FatalOnError(t, e)
	}

	blocked, e := s.Service.Send(r2.Callsign, r1.Callsign, "over", true)
	assert.FatalOnError(t, e)
	assert.False(t, blocked)
	s.Store.Advance(constant.ImmediateWindow + time.Minute)
	allowed, f := s.Service.Send(r2.Callsign, r1.Callsign, "after", true)
	assert.FatalOnError(t, f)
	assert.True(t, allowed)
}
