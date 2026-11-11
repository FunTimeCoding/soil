package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
	"time"
)

func TestImmediateIsRateLimitedToSlowPath(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")

	for range constant.ImmediateLimit {
		result, e := s.Message.Send(
			r2.Callsign,
			[]string{r1.Callsign},
			"burst",
			constant.DeliveryImmediate,
		)
		assert.FatalOnError(t, e)
		assert.Count(t, 0, result.Held)
	}

	result, e := s.Message.Send(
		r2.Callsign,
		[]string{r1.Callsign},
		"over",
		constant.DeliveryImmediate,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, result.Held)
	drained, f := s.Message.DrainImmediateQueue("session-1", r1.Callsign)
	assert.FatalOnError(t, f)
	assert.Count(t, 6, drained)
	remaining := s.Check("session-1")
	assert.Count(
		t,
		1,
		fixture.KindEntries(remaining.Entries, constant.QueueMessage),
	)
}

func TestImmediateRecoversAfterWindow(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")

	for range constant.ImmediateLimit {
		_, e := s.Message.Send(
			r2.Callsign,
			[]string{r1.Callsign},
			"burst",
			constant.DeliveryImmediate,
		)
		assert.FatalOnError(t, e)
	}

	blocked, e := s.Message.Send(
		r2.Callsign,
		[]string{r1.Callsign},
		"over",
		constant.DeliveryImmediate,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, blocked.Held)
	s.Store.Advance(constant.ImmediateWindow + time.Minute)
	allowed, f := s.Message.Send(
		r2.Callsign,
		[]string{r1.Callsign},
		"later",
		constant.DeliveryImmediate,
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 0, allowed.Held)
}
