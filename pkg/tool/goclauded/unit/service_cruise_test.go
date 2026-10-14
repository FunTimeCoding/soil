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

func TestCruiseDeliveryWakesAnArmedIdleSession(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.EndTurn("session-1")
	s.Arm("session-1", constant.CruiseOn, 0)
	assert.True(t, s.Pulse("session-1", constant.DeliveryCruise))
	assert.Integer(t, 1, s.Channel("session-1"))
}

func TestCruiseDeliveryWaitsForThePromptWhileCruiseIsOff(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.EndTurn("session-1")
	assert.False(t, s.Pulse("session-1", constant.DeliveryCruise))
	s.Service.ReleaseCruise()
	assert.Integer(t, 0, s.Channel("session-1"))
	assert.Count(
		t,
		1,
		fixture.EntriesByKind(s.Check("session-1").Entries, constant.QueuePulse),
	)
}

func TestBusySessionHoldsCruiseUntilItsTurnEnds(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.EndTurn("session-1")
	s.Store.Advance(time.Minute)
	s.Check("session-1")
	s.Arm("session-1", constant.CruiseOn, 0)
	assert.False(t, s.Pulse("session-1", constant.DeliveryCruise))
	s.Service.ReleaseCruise()
	assert.Integer(t, 0, s.Channel("session-1"))
	s.EndTurn("session-1")
	s.Service.ReleaseCruise()
	assert.Integer(t, 1, s.Channel("session-1"))
}

func TestPacedCruiseWakesAtMostOncePerPace(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.EndTurn("session-1")
	s.Arm("session-1", constant.CruisePaced, 15*time.Minute)
	assert.True(t, s.Pulse("session-1", constant.DeliveryCruise))
	assert.Integer(t, 1, s.Channel("session-1"))
	assert.False(t, s.Pulse("session-1", constant.DeliveryCruise))
	s.Store.Advance(5 * time.Minute)
	s.Service.ReleaseCruise()
	assert.Integer(t, 0, s.Channel("session-1"))
	s.Store.Advance(15 * time.Minute)
	s.Service.ReleaseCruise()
	assert.Integer(t, 1, s.Channel("session-1"))
}

func TestCruiseReleaseKeepsToTheRateCap(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.EndTurn("session-1")

	for range constant.ImmediateLimit {
		assert.True(t, s.Pulse("session-1", constant.DeliveryImmediate))
	}

	s.Arm("session-1", constant.CruiseOn, 0)
	assert.False(t, s.Pulse("session-1", constant.DeliveryCruise))
	s.Service.ReleaseCruise()
	assert.Integer(t, 6, s.Channel("session-1"))
	s.Store.Advance(constant.ImmediateWindow)
	s.Service.ReleaseCruise()
	assert.Integer(t, 1, s.Channel("session-1"))
}

func TestHeldBacklogReleasesNoFasterThanTheCap(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.EndTurn("session-1")
	s.Store.Advance(time.Minute)
	s.Check("session-1")
	s.Arm("session-1", constant.CruiseOn, 0)

	for range 10 {
		assert.False(t, s.Pulse("session-1", constant.DeliveryCruise))
	}

	s.Store.Advance(2 * constant.ImmediateWindow)
	s.EndTurn("session-1")
	s.Service.ReleaseCruise()
	assert.Integer(t, 6, s.Channel("session-1"))
	s.Service.ReleaseCruise()
	assert.Integer(t, 0, s.Channel("session-1"))
	s.Store.Advance(constant.ImmediateWindow)
	s.Service.ReleaseCruise()
	assert.Integer(t, 4, s.Channel("session-1"))
}

func TestImmediateDeliveryIgnoresCruise(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	assert.True(t, s.Pulse("session-1", constant.DeliveryImmediate))
	assert.Integer(t, 1, s.Channel("session-1"))
}

func TestPacedCruiseNeedsAPace(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	e := s.Service.SetCruise("session-1", constant.CruisePaced, 0)
	assert.True(t, refusal.Is(e))
}

func TestUnknownCruiseModeIsRefused(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	e := s.Service.SetCruise("session-1", constant.Cruise("sometimes"), 0)
	assert.True(t, refusal.Is(e))
}

func TestCruiseIsOffForASessionThatNeverSetIt(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	setting, e := s.Service.CruiseSetting("session-1")
	assert.FatalOnError(t, e)
	assert.String(t, "off", string(setting.Mode))
}
