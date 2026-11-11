package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
	"time"
)

func TestQueuedPulseWaitsOutGrace(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.EndTurn("session-1")
	s.Arm("session-1", "30")
	assert.False(t, s.Pulse("session-1", constant.DeliveryQueued))
	s.Store.Advance(29 * time.Minute)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 0, s.Channel("session-1"))
	s.Store.Advance(time.Minute)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 1, s.Channel("session-1"))
}

func TestOneWakeDeliversEverythingWaiting(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.EndTurn("session-1")
	s.Arm("session-1", "30")
	s.Pulse("session-1", constant.DeliveryQueued)
	s.Send(r2.Callsign, r1.Callsign, "named")
	s.Send(r2.Callsign, "", "to everyone")
	_, e := s.Message.SendNotification(
		r1.Callsign,
		"mattermost",
		"notice",
		constant.DeliveryQueued,
	)
	assert.FatalOnError(t, e)
	s.Store.Advance(30 * time.Minute)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 4, s.Channel("session-1"))
	setting, f := s.Cruise.Setting("session-1")
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, setting.Woken)
}

func TestRosterActivityNeverPops(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.EndTurn("session-1")
	s.Arm("session-1", "30")
	s.Announce("session-2", r2.Callsign, "some work", "")
	s.Store.Advance(30 * time.Minute)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 0, s.Channel("session-1"))
	assert.Count(
		t,
		1,
		fixture.KindEntries(
			s.Check("session-1").Entries,
			constant.QueueSessionAnnounce,
		),
	)
}

func TestCruiseOffWaitsForPrompt(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.EndTurn("session-1")
	s.Arm("session-1", constant.CruiseOff)
	s.Pulse("session-1", constant.DeliveryQueued)
	s.Store.Advance(30 * time.Minute)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 0, s.Channel("session-1"))
	assert.Count(
		t,
		1,
		fixture.KindEntries(s.Check("session-1").Entries, constant.QueuePulse),
	)
}

func TestGraceCountsFromTurnEnd(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.EndTurn("session-1")
	s.Arm("session-1", "30")
	s.Store.Advance(40 * time.Minute)
	s.Check("session-1")
	s.Pulse("session-1", constant.DeliveryQueued)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 0, s.Channel("session-1"))
	s.EndTurn("session-1")
	s.Store.Advance(29 * time.Minute)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 0, s.Channel("session-1"))
	s.Store.Advance(time.Minute)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 1, s.Channel("session-1"))
}

func TestReleaseCountsOnceAgainstCap(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.EndTurn("session-1")
	s.Arm("session-1", "30")

	for range 10 {
		s.Pulse("session-1", constant.DeliveryQueued)
	}

	s.Store.Advance(30 * time.Minute)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 10, s.Channel("session-1"))

	for range constant.ImmediateLimit - 1 {
		assert.True(t, s.Pulse("session-1", constant.DeliveryImmediate))
	}

	assert.False(t, s.Pulse("session-1", constant.DeliveryImmediate))
}

func TestReleaseWaitsWhileCapIsSpent(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.EndTurn("session-1")
	s.Arm("session-1", "30")
	s.Store.Advance(30 * time.Minute)

	for range constant.ImmediateLimit {
		assert.True(t, s.Pulse("session-1", constant.DeliveryImmediate))
	}

	s.Pulse("session-1", constant.DeliveryQueued)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 6, s.Channel("session-1"))
	s.Store.Advance(constant.ImmediateWindow)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 1, s.Channel("session-1"))
}

func TestImmediateDeliveryIgnoresCruise(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	assert.True(t, s.Pulse("session-1", constant.DeliveryImmediate))
	assert.Integer(t, 1, s.Channel("session-1"))
}

func TestCruiseIsOffByDefaultWithoutChannelPromise(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.EndTurn("session-1")
	s.Pulse("session-1", constant.DeliveryQueued)
	s.Store.Advance(30 * time.Minute)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 0, s.Channel("session-1"))
	setting, e := s.Cruise.Setting("session-1")
	assert.FatalOnError(t, e)
	assert.Nil(t, setting.Grace)
	assert.Duration(t, 0, s.Cruise.Grace(setting))
}

func TestDefaultGraceWakesSessionWithoutItsOwn(t *testing.T) {
	s := service_tester.New(t)
	s.Cruise.WithDefaultGrace(constant.DefaultCruiseGrace)
	s.Check("session-1")
	s.EndTurn("session-1")
	s.Pulse("session-1", constant.DeliveryQueued)
	s.Store.Advance(30 * time.Minute)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 1, s.Channel("session-1"))
	setting, e := s.Cruise.Setting("session-1")
	assert.FatalOnError(t, e)
	assert.Nil(t, setting.Grace)
	assert.Integer(t, 1, setting.Woken)
}

func TestOwnGraceOverridesDefault(t *testing.T) {
	s := service_tester.New(t)
	s.Cruise.WithDefaultGrace(constant.DefaultCruiseGrace)
	s.Check("session-1")
	s.EndTurn("session-1")
	s.Arm("session-1", "10")
	s.Pulse("session-1", constant.DeliveryQueued)
	s.Store.Advance(10 * time.Minute)
	assert.Count(t, 0, s.Cruise.Release())
	assert.Integer(t, 1, s.Channel("session-1"))
}

func TestDefaultClearsOwnGrace(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.Arm("session-1", "10")
	s.Arm("session-1", constant.CruiseDefault)
	setting, e := s.Cruise.Setting("session-1")
	assert.FatalOnError(t, e)
	assert.Nil(t, setting.Grace)
}

func TestCruiseGraceTakesOneToFiftyNineMinutes(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	assert.FatalOnError(t, s.Cruise.Set("session-1", "1"))
	assert.FatalOnError(t, s.Cruise.Set("session-1", "59"))

	for _, refused := range []string{"0", "60", "1.5", "sometimes", ""} {
		assert.True(t, refusal.Is(s.Cruise.Set("session-1", refused)))
	}
}
