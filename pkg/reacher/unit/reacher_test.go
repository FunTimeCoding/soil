package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/reacher"
	"testing"
	"time"
)

func TestReacherEdgesOnlyOnTransition(t *testing.T) {
	c := newClock()
	r := reacher.New().WithClock(c.Now)
	down := r.Fail("sentry", constant.Refused)
	assert.NotNil(t, down)
	assert.True(t, down.Down)
	assert.String(t, "sentry", down.Host)
	assert.String(t, "connection refused", down.Reason)
	assert.Nil(t, r.Fail("sentry", constant.Refused))
	assert.Nil(t, r.Fail("sentry", constant.TimedOut))
	c.Advance(130 * time.Second)
	up := r.Succeed("sentry")
	assert.NotNil(t, up)
	assert.False(t, up.Down)
	assert.Duration(t, 130*time.Second, up.Duration)
	assert.Nil(t, r.Succeed("sentry"))
}

func TestReacherKeepsFirstReasonThroughOutage(t *testing.T) {
	r := reacher.New()
	r.Fail("sentry", constant.Refused)
	r.Fail("sentry", constant.TimedOut)
	down, h := r.Down("sentry")
	assert.True(t, down)
	assert.String(t, "connection refused", h.Reason)
}

func TestReacherHostsAreIndependent(t *testing.T) {
	r := reacher.New()
	assert.NotNil(t, r.Fail("sentry", constant.Refused))
	assert.NotNil(t, r.Fail("netbox", constant.NoRoute))
	assert.NotNil(t, r.Succeed("sentry"))
	down, _ := r.Down("netbox")
	assert.True(t, down)
	down, _ = r.Down("sentry")
	assert.False(t, down)
}

func TestReacherSuccessOnUnknownHostIsSilent(t *testing.T) {
	r := reacher.New()
	assert.Nil(t, r.Succeed("sentry"))
	down, h := r.Down("sentry")
	assert.False(t, down)
	assert.Nil(t, h)
}

func TestReacherDownReturnsACopy(t *testing.T) {
	r := reacher.New()
	r.Fail("sentry", constant.Refused)
	_, h := r.Down("sentry")
	h.Down = false
	down, _ := r.Down("sentry")
	assert.True(t, down)
}
