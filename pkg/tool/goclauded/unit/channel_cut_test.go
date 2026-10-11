package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/channel_tester"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

func TestChannelDeliveryCutsALongMessage(t *testing.T) {
	s := base.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	x, cancel := context.WithCancel(t.Context())
	defer cancel()
	result := channel_tester.Deliver(t, s, x, r1.Callsign)
	s.AwaitSubscribers(1)
	s.SendImmediate(r2.Callsign, r1.Callsign, strings.Repeat("a", 5000))

	select {
	case entries := <-result:
		assert.Count(t, 1, entries)
		assert.StringContains(t, "…cut here - ", entries[0].Body)
		assert.StringContains(t, "read it with read_message", entries[0].Body)
		assert.True(t, utf8.RuneCountInString(entries[0].Body) <= 2000)
	case <-time.After(2 * time.Second):
		t.Fatal("delivery did not wake on the immediate entry")
	}
}

func TestChannelDeliveryKeepsAShortMessageWhole(t *testing.T) {
	s := base.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	x, cancel := context.WithCancel(t.Context())
	defer cancel()
	result := channel_tester.Deliver(t, s, x, r1.Callsign)
	s.AwaitSubscribers(1)
	s.SendImmediate(r2.Callsign, r1.Callsign, "short")

	select {
	case entries := <-result:
		assert.Count(t, 1, entries)
		assert.String(t, join.Empty(r2.Callsign, ": short"), entries[0].Body)
	case <-time.After(2 * time.Second):
		t.Fatal("delivery did not wake on the immediate entry")
	}
}
