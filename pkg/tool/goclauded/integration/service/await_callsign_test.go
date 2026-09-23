package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"testing"
	"time"
)

func TestAwaitCallsignHoldsUntilTheSessionTakesAPrompt(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("session-1")
	assert.String(
		t,
		"",
		s.AwaitCallsign(fixture.Cancelled(), "session-1", s.Store.Clock()()),
	)
}

func TestAwaitCallsignReturnsOnAPromptAfterTheChannelStarted(t *testing.T) {
	s := service_tester.New(t)
	since := s.Store.Clock()().Add(-time.Second)
	r := s.Check("session-1")
	assert.String(
		t,
		r.Callsign,
		s.AwaitCallsign(fixture.Cancelled(), "session-1", since),
	)
}

func TestAwaitCallsignIgnoresAPromptOlderThanTheChannel(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.Store.Advance(time.Minute)
	assert.String(
		t,
		"",
		s.AwaitCallsign(fixture.Cancelled(), "session-1", s.Store.Clock()()),
	)
}

func TestAwaitCallsignNeedsNoAnnounce(t *testing.T) {
	s := service_tester.New(t)
	since := s.Store.Clock()().Add(-time.Second)
	r := s.Check("session-1")
	assert.String(t, "", s.Store.GetSession("session-1").Topic)
	assert.String(
		t,
		r.Callsign,
		s.AwaitCallsign(fixture.Cancelled(), "session-1", since),
	)
}

func TestAwaitCallsignWakesOnTheNextPrompt(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Store.Advance(time.Minute)
	since := s.Store.Clock()()
	result := make(chan string, 1)
	go func() {
		result <- s.AwaitCallsign(context.Background(), "session-1", since)
	}()

	select {
	case <-result:
		t.Fatal("callsign returned before the next prompt")
	case <-time.After(50 * time.Millisecond):
	}

	s.Store.Advance(time.Minute)
	s.Check("session-1")

	select {
	case callsign := <-result:
		assert.String(t, r.Callsign, callsign)
	case <-time.After(time.Second):
		t.Fatal("await did not wake on the prompt")
	}
}

func TestAwaitCallsignReleasesWhenTheCallerGoesAway(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("session-1")
	x, cancel := context.WithCancel(context.Background())
	result := make(chan string, 1)
	go func() {
		result <- s.AwaitCallsign(x, "session-1", s.Store.Clock()())
	}()
	time.Sleep(50 * time.Millisecond)
	cancel()

	select {
	case callsign := <-result:
		assert.String(t, "", callsign)
	case <-time.After(time.Second):
		t.Fatal("await did not release on cancellation")
	}
}

func TestAwaitCallsignReleasesWhenTheServiceStops(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("session-1")
	result := make(chan string, 1)
	go func() {
		result <- s.AwaitCallsign(
			context.Background(),
			"session-1",
			s.Store.Clock()(),
		)
	}()
	time.Sleep(50 * time.Millisecond)
	s.Service.Stop()

	select {
	case callsign := <-result:
		assert.String(t, "", callsign)
	case <-time.After(time.Second):
		t.Fatal("await did not release when the service stopped")
	}
}
