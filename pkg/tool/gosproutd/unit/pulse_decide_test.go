package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/pulse"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/pulse_tester"
	"testing"
	"time"
)

func TestNothingUnpulsedDrops(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.Unpulsed = 0
	assert.String(t, "drop", string(pulse.Decide(r, now)))
}

func TestClosedSessionDrops(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.Closed = true
	assert.String(t, "drop", string(pulse.Decide(r, now)))
}

func TestFreshAnswerHoldsForTheQuietWindow(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.NewestAt = now.Add(-time.Minute)
	assert.String(t, "hold", string(pulse.Decide(r, now)))
}

func TestNewAnswersCannotHoldPastTheCeiling(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.NewestAt = now.Add(-time.Second)
	r.OldestAt = now.Add(-constant.MaximumHold - time.Minute)
	assert.String(t, "immediate", string(pulse.Decide(r, now)))
}

func TestCruiseOffNeverPromotes(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.Mode = constant.CruiseOff
	assert.String(t, "queue", string(pulse.Decide(r, now)))
}

func TestIdleSessionInsideTheWakeWindowIsWoken(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.IdleSince = now.Add(-30 * time.Minute)
	assert.String(t, "immediate", string(pulse.Decide(r, now)))
}

func TestIdleSessionPastTheWakeWindowIsNotWoken(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.IdleSince = now.Add(-constant.WakeWindow - time.Minute)
	assert.String(t, "queue", string(pulse.Decide(r, now)))
}

func TestWorkingSessionIsNeverWoken(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.Idle = false
	assert.String(t, "queue", string(pulse.Decide(r, now)))
}

func TestUnknownStateIsNeverWoken(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.StateKnown = false
	assert.String(t, "queue", string(pulse.Decide(r, now)))
}

func TestPacedModeWaitsOutTheGap(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.Mode = constant.CruisePaced
	r.Pace = 15 * time.Minute
	r.PromotedAt = now.Add(-5 * time.Minute)
	assert.String(t, "hold", string(pulse.Decide(r, now)))
}

func TestPacedModePromotesAfterTheGap(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.Mode = constant.CruisePaced
	r.Pace = 15 * time.Minute
	r.PromotedAt = now.Add(-20 * time.Minute)
	assert.String(t, "immediate", string(pulse.Decide(r, now)))
}

func TestPacedModeFirstPromotionDoesNotWait(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.Mode = constant.CruisePaced
	r.Pace = 30 * time.Minute
	assert.String(t, "immediate", string(pulse.Decide(r, now)))
}

func TestPacedGapYieldsWhenSheStepsInHerself(t *testing.T) {
	now := time.Now()
	r := pulse_tester.Settled(now)
	r.Mode = constant.CruisePaced
	r.Pace = 15 * time.Minute
	r.PromotedAt = now.Add(-20 * time.Minute)
	r.Idle = false
	assert.String(t, "queue", string(pulse.Decide(r, now)))
}
