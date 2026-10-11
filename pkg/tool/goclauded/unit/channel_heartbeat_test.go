package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclaude"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/channel_tester"
	"testing"
	"time"
)

func TestQuietChannelNeverStalls(t *testing.T) {
	s := base.NewWithHold(t, constant.FixtureChannelHold)
	r := s.Check("session-1")
	c := s.RESTClient(t)
	server, k := channel_tester.Sink()
	failures := 0

	for range 6 {
		failures = goclaude.PollChannel(c, server, r.Callsign, failures)
	}

	assert.Integer(t, 0, failures)
	assert.Count(t, 0, k.Kind("stalled"))
	assert.Count(t, 0, k.Events())
}

func TestQuietChannelStillDeliversAfterHolds(t *testing.T) {
	s := base.NewWithHold(t, constant.FixtureChannelHold)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	c := s.RESTClient(t)
	server, k := channel_tester.Sink()
	failures := 0

	for range 3 {
		failures = goclaude.PollChannel(c, server, r1.Callsign, failures)
	}

	s.SendImmediate(r2.Callsign, r1.Callsign, "after the quiet")
	failures = goclaude.PollChannel(c, server, r1.Callsign, failures)
	assert.Integer(t, 0, failures)
	messages := k.Kind(constant.QueueMessage)
	assert.Count(t, 1, messages)
	assert.String(
		t,
		join.Empty(r2.Callsign, ": after the quiet"),
		messages[0].Content,
	)
}

func TestCallsignWaitOutlastsReleasedHolds(t *testing.T) {
	s := base.NewWithHold(t, constant.FixtureChannelHold)
	r := s.Check("session-1")
	s.Store.Advance(time.Minute)
	since := s.Store.Clock()()
	c := s.RESTClient(t)
	result := make(chan string, 1)
	go func() {
		result <- goclaude.AwaitCallsign(
			c,
			"session-1",
			since,
			time.Millisecond,
		)
	}()

	select {
	case <-result:
		t.Fatal("callsign returned before the next prompt")
	case <-time.After(10 * constant.FixtureChannelHold):
	}

	s.Store.Advance(time.Minute)
	s.Check("session-1")

	select {
	case name := <-result:
		assert.String(t, r.Callsign, name)
	case <-time.After(2 * time.Second):
		t.Fatal("callsign did not arrive after the prompt")
	}
}
