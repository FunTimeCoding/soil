package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"testing"
	"time"
)

func TestAwaitImmediateQueueReturnsEntriesAlreadyWaiting(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.SendImmediate(r2.Callsign, r1.Callsign, "immediate message")
	drained := s.AwaitImmediateQueue(
		fixture.Cancelled(),
		"session-1",
		r1.Callsign,
	)
	assert.Count(t, 1, drained)
	assert.String(t, "message", drained[0].Kind)
}

func TestAwaitImmediateQueueReturnsNothingWhenTheQueueIsEmpty(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	assert.Count(
		t,
		0,
		s.AwaitImmediateQueue(fixture.Cancelled(), "session-1", r1.Callsign),
	)
}

func TestAwaitImmediateQueueWakesWhenAnEntryArrivesDuringTheWait(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	result := make(chan []queue.Entry, 1)
	go func() {
		result <- s.AwaitImmediateQueue(
			context.Background(),
			"session-1",
			r1.Callsign,
		)
	}()
	time.Sleep(50 * time.Millisecond)
	s.SendImmediate(r2.Callsign, r1.Callsign, "immediate message")

	select {
	case drained := <-result:
		assert.Count(t, 1, drained)
		assert.String(t, "message", drained[0].Kind)
	case <-time.After(time.Second):
		t.Fatal("await did not wake on the immediate entry")
	}
}

func TestAwaitImmediateQueueReleasesWhenTheServiceStops(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	result := make(chan []queue.Entry, 1)
	go func() {
		result <- s.AwaitImmediateQueue(
			context.Background(),
			"session-1",
			r1.Callsign,
		)
	}()
	time.Sleep(50 * time.Millisecond)
	s.Service.Stop()

	select {
	case drained := <-result:
		assert.Count(t, 0, drained)
	case <-time.After(time.Second):
		t.Fatal("await did not release when the service stopped")
	}
}
