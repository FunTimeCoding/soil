package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/peer/outage_tracker"
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"testing"
	"time"
)

func TestOutageTrackerEdgesOnlyOnTransition(t *testing.T) {
	c := clock_tester.New()
	r := outage_tracker.New(memory.New()).WithClock(c.Now)
	down := r.Fail("sentry", errorsConstant.Refused)
	assert.NotNil(t, down)
	assert.True(t, down.Down)
	assert.String(t, "sentry", down.Host)
	assert.String(t, "connection refused", down.Reason)
	assert.Nil(t, r.Fail("sentry", errorsConstant.Refused))
	assert.Nil(t, r.Fail("sentry", errorsConstant.TimedOut))
	c.Advance(130 * time.Second)
	up := r.Succeed("sentry")
	assert.NotNil(t, up)
	assert.False(t, up.Down)
	assert.Duration(t, 130*time.Second, up.Duration)
	assert.Nil(t, r.Succeed("sentry"))
}

func TestOutageTrackerKeepsFirstReasonThroughOutage(t *testing.T) {
	r := outage_tracker.New(memory.New())
	r.Fail("sentry", errorsConstant.Refused)
	r.Fail("sentry", errorsConstant.TimedOut)
	down, h := r.Down("sentry")
	assert.True(t, down)
	assert.String(t, "connection refused", h.Reason)
}

func TestOutageTrackerHostsAreIndependent(t *testing.T) {
	r := outage_tracker.New(memory.New())
	assert.NotNil(t, r.Fail("sentry", errorsConstant.Refused))
	assert.NotNil(t, r.Fail("netbox", errorsConstant.NoRoute))
	assert.NotNil(t, r.Succeed("sentry"))
	down, _ := r.Down("netbox")
	assert.True(t, down)
	down, _ = r.Down("sentry")
	assert.False(t, down)
}

func TestOutageTrackerSuccessOnUnknownHostIsSilent(t *testing.T) {
	r := outage_tracker.New(memory.New())
	assert.Nil(t, r.Succeed("sentry"))
	down, h := r.Down("sentry")
	assert.False(t, down)
	assert.Nil(t, h)
}

func TestOutageTrackerDownReturnsCopy(t *testing.T) {
	r := outage_tracker.New(memory.New())
	r.Fail("sentry", errorsConstant.Refused)
	_, h := r.Down("sentry")
	h.Down = false
	down, _ := r.Down("sentry")
	assert.True(t, down)
}
