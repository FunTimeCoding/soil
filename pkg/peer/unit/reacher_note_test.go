package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/peer/unit/fake_clock"
	"testing"
	"time"
)

func TestANotedOutageLogsItsEdgesButNeverReports(t *testing.T) {
	c := fake_clock.New()
	m := memory.New()
	r := reacher.New(m).WithClock(c.Now)
	down, e := r.Note("ilo", refusedConnection())
	assert.Nil(t, e)
	assert.NotNil(t, down)
	c.Advance(time.Hour)
	again, f := r.Note("ilo", refusedConnection())
	assert.Nil(t, f)
	assert.Nil(t, again)
	assert.Count(t, 0, m.Events())
	up, g := r.Note("ilo", nil)
	assert.Nil(t, g)
	assert.NotNil(t, up)
}

func TestANotedRefusalComesBackOncePerReason(t *testing.T) {
	m := memory.New()
	r := reacher.New(m)
	_, e := r.Note("ilo", errors.Format("iLO error", "0x003C"))
	assert.Error(t, e)
	_, f := r.Note("ilo", errors.Format("iLO error", "0x003C"))
	assert.Nil(t, f)
	_, g := r.Note("ilo", errors.Format("iLO error", "0x0004"))
	assert.Error(t, g)
	down, _ := r.Down("ilo")
	assert.False(t, down)
	assert.Count(t, 0, m.Events())
}
