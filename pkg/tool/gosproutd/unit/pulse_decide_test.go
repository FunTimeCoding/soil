package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/pulse"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/pulse_tester"
	"testing"
	"time"
)

func TestSettledAnswersAreSent(t *testing.T) {
	now := time.Now()
	assert.String(
		t,
		"send",
		string(pulse.Decide(pulse_tester.Settled(now), now)),
	)
}

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
	assert.String(t, "send", string(pulse.Decide(r, now)))
}
