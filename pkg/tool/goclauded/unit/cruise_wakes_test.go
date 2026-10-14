package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/cruise"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/cruise_tester"
	"testing"
	"time"
)

func TestArmedIdleSessionIsWoken(t *testing.T) {
	now := time.Now()
	assert.True(t, cruise.Wakes(cruise_tester.Settled(now), now))
}

func TestCruiseOffNeverWakes(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.Mode = constant.CruiseOff
	assert.False(t, cruise.Wakes(r, now))
}

func TestClosedSessionIsNeverWoken(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.Closed = true
	assert.False(t, cruise.Wakes(r, now))
}

func TestBusySessionIsNotWokenMidTurn(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.Idle = false
	assert.False(t, cruise.Wakes(r, now))
}

func TestUnknownStateIsNeverWoken(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.StateKnown = false
	assert.False(t, cruise.Wakes(r, now))
}

func TestIdleSessionInsideTheWakeWindowIsWoken(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.IdleSince = now.Add(-30 * time.Minute)
	assert.True(t, cruise.Wakes(r, now))
}

func TestIdleSessionPastTheWakeWindowIsNotWoken(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.IdleSince = now.Add(-constant.WakeWindow - time.Minute)
	assert.False(t, cruise.Wakes(r, now))
}

func TestPacedModeWaitsOutTheGap(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.Mode = constant.CruisePaced
	r.Pace = 15 * time.Minute
	r.PromotedAt = now.Add(-5 * time.Minute)
	assert.False(t, cruise.Wakes(r, now))
}

func TestPacedModeWakesAfterTheGap(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.Mode = constant.CruisePaced
	r.Pace = 15 * time.Minute
	r.PromotedAt = now.Add(-20 * time.Minute)
	assert.True(t, cruise.Wakes(r, now))
}

func TestPacedModeFirstWakeDoesNotWait(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.Mode = constant.CruisePaced
	r.Pace = 30 * time.Minute
	assert.True(t, cruise.Wakes(r, now))
}

func TestPacedGapYieldsWhenSheStepsInHerself(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.Mode = constant.CruisePaced
	r.Pace = 15 * time.Minute
	r.PromotedAt = now.Add(-20 * time.Minute)
	r.Idle = false
	assert.False(t, cruise.Wakes(r, now))
}
