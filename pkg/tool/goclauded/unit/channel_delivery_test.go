package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/channel_tester"
	"testing"
	"time"
)

func TestChannelDeliveryHoldsWhileNothingIsPending(t *testing.T) {
	s := base.New(t)
	r := s.Check("session-1")
	x, cancel := context.WithCancel(t.Context())
	defer cancel()
	result := channel_tester.Deliver(t, s, x, r.Callsign)

	select {
	case <-result:
		t.Fatal("delivery returned while the queue was empty")
	case <-time.After(250 * time.Millisecond):
	}
}

func TestChannelDeliveryReturnsWhenAnImmediateEntryArrives(t *testing.T) {
	s := base.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	x, cancel := context.WithCancel(t.Context())
	defer cancel()
	result := channel_tester.Deliver(t, s, x, r1.Callsign)
	s.AwaitSubscribers(1)
	s.SendImmediate(r2.Callsign, r1.Callsign, "immediate message")

	select {
	case entries := <-result:
		assert.Count(t, 1, entries)
		assert.String(t, "message", entries[0].Kind)
	case <-time.After(2 * time.Second):
		t.Fatal("delivery did not wake on the immediate entry")
	}
}

func TestChannelDeliveryIgnoresADeferredEntry(t *testing.T) {
	s := base.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	x, cancel := context.WithCancel(t.Context())
	defer cancel()
	result := channel_tester.Deliver(t, s, x, r1.Callsign)
	s.AwaitSubscribers(1)
	s.Send(r2.Callsign, r1.Callsign, "deferred message")

	select {
	case <-result:
		t.Fatal("delivery returned for an entry the hook owns")
	case <-time.After(500 * time.Millisecond):
	}
}
