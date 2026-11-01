package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/peer/constant"
	"github.com/funtimecoding/soil/pkg/peer/outage"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/peer/unit/fake_clock"
	"os"
	"testing"
	"time"
)

func TestAShortOutageStaysOutOfSentry(t *testing.T) {
	c := fake_clock.New()
	m := memory.New()
	r := reacher.New(m).WithClock(c.Now)
	r.Observe("sentry", refusedConnection())
	c.Advance(4 * time.Minute)
	r.Observe("sentry", refusedConnection())
	r.Observe("sentry", nil)
	assert.Count(t, 0, m.Events())
}

func TestALastingOutageIsReportedOncePerEpisode(t *testing.T) {
	c := fake_clock.New()
	m := memory.New()
	r := reacher.New(m).WithClock(c.Now)
	r.Observe("sentry", refusedConnection())
	c.Advance(5 * time.Minute)
	r.Observe("sentry", refusedConnection())
	c.Advance(time.Hour)
	r.Observe("sentry", refusedConnection())
	assert.Count(t, 1, m.Events())
	r.Observe("sentry", nil)
	r.Observe("sentry", refusedConnection())
	c.Advance(5 * time.Minute)
	r.Observe("sentry", refusedConnection())
	assert.Count(t, 2, m.Events())
}

func TestAnOutageTheCallerNamesLastsLikeAnyOther(t *testing.T) {
	c := fake_clock.New()
	m := memory.New()
	r := reacher.New(m).WithClock(c.Now)
	assert.NotNil(
		t,
		r.Outage("serial", constant.NotConnectedReason, os.ErrNotExist),
	)
	c.Advance(5 * time.Minute)
	assert.Nil(
		t,
		r.Outage("serial", constant.NotConnectedReason, os.ErrNotExist),
	)
	assert.Nil(
		t,
		r.Outage("serial", constant.NotConnectedReason, os.ErrNotExist),
	)
	assert.Count(t, 1, m.Events())
	assert.String(
		t,
		"serial down for 5m0s: file does not exist",
		m.Events()[0].Error.Error(),
	)
}

func TestALastingOutageReportsTheTransportFailureWithItsEpisode(t *testing.T) {
	c := fake_clock.New()
	m := memory.New()
	r := reacher.New(m).WithClock(c.Now)
	r.Observe("sentry", refusedConnection())
	c.Advance(6 * time.Minute)
	r.Observe("sentry", refusedConnection())
	assert.Count(t, 1, m.Events())
	e := m.Events()[0].Error
	assert.String(
		t,
		"sentry down for 6m0s: dial tcp: connect: connection refused",
		e.Error(),
	)
	assert.NotNil(t, connection.Classify(e))
	o, okay := e.(*outage.OutageError)
	assert.True(t, okay)
	key, fields := o.ErrorContext()
	assert.String(t, "reachability", key)
	assert.Any(
		t,
		map[string]any{
			"host":     "sentry",
			"state":    "down",
			"since":    "2026-10-06 11:23:00",
			"duration": "6m0s",
		},
		fields,
	)
}
