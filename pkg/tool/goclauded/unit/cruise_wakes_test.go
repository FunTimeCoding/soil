package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/cruise_tester"
	"testing"
	"time"
)

func TestIdleSessionPastGraceIsWoken(t *testing.T) {
	now := time.Now()
	assert.True(t, cruise_tester.Settled(now).Wakes(now))
}

func TestCruiseOffNeverWakes(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.Grace = 0
	assert.False(t, r.Wakes(now))
}

func TestClosedSessionIsNeverWoken(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.Closed = true
	assert.False(t, r.Wakes(now))
}

func TestBusySessionIsNotWokenMidTurn(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.Idle = false
	assert.False(t, r.Wakes(now))
}

func TestUnknownStateIsNeverWoken(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.StateKnown = false
	assert.False(t, r.Wakes(now))
}

func TestIdleSessionInsideGraceIsNotWoken(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.IdleSince = now.Add(-29 * time.Minute)
	assert.False(t, r.Wakes(now))
}

func TestIdleSessionAtGraceIsWoken(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.IdleSince = now.Add(-30 * time.Minute)
	assert.True(t, r.Wakes(now))
}

func TestIdleSessionPastWakeWindowIsNotWoken(t *testing.T) {
	now := time.Now()
	r := cruise_tester.Settled(now)
	r.IdleSince = now.Add(-constant.WakeWindow)
	assert.False(t, r.Wakes(now))
}
