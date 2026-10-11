package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestDeliveryKeepsShortMessagesWhole(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Send(r2.Callsign, r1.Callsign, "hello")
	assert.String(
		t,
		fmt.Sprintf("Messages:\n  %s: hello", r2.Callsign),
		s.Check("session-1").Context,
	)
}

func TestDeliveryCutsLongMessage(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Send(r2.Callsign, r1.Callsign, strings.Repeat("a", 5000))
	text := s.Check("session-1").Context
	lines := strings.Split(text, "\n")
	assert.Count(t, 3, lines)
	assert.String(
		t,
		"  This delivery holds about 2,000 characters; longer messages are cut or waiting, and read_message returns them whole.",
		lines[1],
	)
	assert.StringContains(t, "…cut here - ", lines[2])
	assert.StringContains(t, "read it with read_message", lines[2])
	assert.True(t, utf8.RuneCountInString(text) <= 2000)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestDeliveryFillsBudgetAfterCut(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Send(r2.Callsign, r1.Callsign, "first")
	s.Send(r2.Callsign, r1.Callsign, strings.Repeat("b", 5000))
	s.Send(r2.Callsign, r1.Callsign, "third")
	lines := strings.Split(s.Check("session-1").Context, "\n")
	assert.Count(t, 5, lines)
	assert.String(t, fmt.Sprintf("  %s: first", r2.Callsign), lines[2])
	assert.StringContains(t, "…cut here - ", lines[3])
	assert.String(t, fmt.Sprintf("  %s: third", r2.Callsign), lines[4])
}

func TestDeliveryWaitsBelowCutFloor(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")

	for range 4 {
		s.Send(r2.Callsign, r1.Callsign, strings.Repeat("c", 3000))
	}

	text := s.Check("session-1").Context
	lines := strings.Split(text, "\n")
	assert.Count(t, 6, lines)
	assert.StringContains(t, "…cut here - ", lines[2])
	assert.StringContains(t, "(3,000 characters) is waiting", lines[5])
	assert.True(t, utf8.RuneCountInString(text) <= 2000)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestDeliveryTrimsMemoryActivity(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")

	for i := range 70 {
		assert.FatalOnError(
			t,
			s.Store.Store.PushQueue(
				"session-1",
				r1.Callsign,
				constant.QueueMemoryUpdate,
				fmt.Sprintf("memory number %02d updated by Ellis", i),
			),
		)
	}

	s.Send(r2.Callsign, r1.Callsign, "after the memories")
	text := s.Check("session-1").Context
	assert.StringContains(t, "  Memory changes not shown: ", text)
	assert.StringContains(
		t,
		fmt.Sprintf("  %s: after the memories", r2.Callsign),
		text,
	)
	assert.True(t, utf8.RuneCountInString(text) <= 2000)
}

func TestDeliveryOverBudgetReachesSentry(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")

	for i := range 30 {
		assert.FatalOnError(
			t,
			s.Store.Store.PushQueue(
				"session-1",
				r1.Callsign,
				constant.QueueSessionUpdate,
				fmt.Sprintf(
					"Session%02d updated scope: %s",
					i,
					strings.Repeat("d", 80),
				),
			),
		)
	}

	s.Check("session-1")
	events := s.Reporter.Events()
	assert.Count(t, 1, events)
	assert.String(t, "hook delivery over budget", events[0].Error.Error())
}

func TestChannelCutsLongMessage(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.SendImmediate(r2.Callsign, r1.Callsign, strings.Repeat("e", 5000))
	drained, e := s.Service.DrainImmediateQueue("session-1", r1.Callsign)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, drained)
	assert.StringContains(t, "…cut here - ", drained[0].Body)
	assert.True(t, utf8.RuneCountInString(drained[0].Body) <= 2000)
}

func TestNotificationOverLimitIsRefused(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	_, e := s.Service.SendNotification(
		r1.Callsign,
		"mattermost",
		strings.Repeat("f", 501),
		false,
	)
	assert.True(t, refusal.Is(e))
	assert.String(
		t,
		"notification is 501 characters; summarize it to 500 or fewer",
		e.Error(),
	)
}

func TestInboundPulseOverLimitIsRefused(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	_, e := s.Service.SendPulse(
		"session-1",
		"",
		strings.Repeat("g", 501),
		false,
	)
	assert.True(t, refusal.Is(e))
}

func TestOutboundPulseIsNotCapped(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	_, e := s.Service.SendPulse(
		"session-1",
		r1.Callsign,
		strings.Repeat("h", 501),
		false,
	)
	assert.FatalOnError(t, e)
}
