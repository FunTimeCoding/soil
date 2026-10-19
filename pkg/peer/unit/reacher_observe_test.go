package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
	"syscall"
	"testing"
)

func TestObserveClassifiesTransportFailure(t *testing.T) {
	r := reacher.New(memory.New())
	e := r.Observe("sentry", syscall.ECONNREFUSED)
	assert.NotNil(t, e)
	assert.True(t, e.Down)
	assert.String(t, "connection refused", e.Reason)
	assert.Nil(t, r.Observe("sentry", syscall.ECONNREFUSED))
	up := r.Observe("sentry", nil)
	assert.NotNil(t, up)
	assert.False(t, up.Down)
}

func TestObserveIgnoresNonTransportErrors(t *testing.T) {
	r := reacher.New(memory.New())
	assert.Nil(t, r.Observe("sentry", not_found.New("issue", "GO-1")))
	down, _ := r.Down("sentry")
	assert.False(t, down)
	r.Fail("sentry", constant.Refused)
	assert.Nil(t, r.Observe("sentry", not_found.New("issue", "GO-1")))
	down, _ = r.Down("sentry")
	assert.True(t, down)
}

func TestTriageHandsBackWhatIsNotAnOutage(t *testing.T) {
	r := reacher.New(memory.New())
	down, real := r.Triage("sentry", syscall.ECONNREFUSED)
	assert.NotNil(t, down)
	assert.Nil(t, real)
	again, quiet := r.Triage("sentry", syscall.ECONNREFUSED)
	assert.Nil(t, again)
	assert.Nil(t, quiet)
	refused := not_found.New("issue", "GO-1")
	none, back := r.Triage("sentry", refused)
	assert.Nil(t, none)
	assert.True(t, back == refused)
	up, clean := r.Triage("sentry", nil)
	assert.NotNil(t, up)
	assert.False(t, up.Down)
	assert.Nil(t, clean)
}

func TestHandleLogsAnOutageOnceAndEachRefusalOncePerReason(t *testing.T) {
	r := reacher.New(memory.New())
	var crossed []*edge.Edge
	var refused []error
	handle := func(e error) {
		r.Handle(
			"sentry",
			e,
			func(x *edge.Edge) { crossed = append(crossed, x) },
			func(f error) { refused = append(refused, f) },
		)
	}
	handle(syscall.ECONNREFUSED)
	handle(syscall.ECONNREFUSED)
	handle(not_found.New("issue", "GO-1"))
	handle(not_found.New("issue", "GO-1"))
	handle(not_found.New("issue", "GO-2"))
	handle(nil)
	assert.Count(t, 2, crossed)
	assert.True(t, crossed[0].Down)
	assert.False(t, crossed[1].Down)
	assert.Count(t, 2, refused)
	assert.StringContains(t, "GO-1", refused[0].Error())
	assert.StringContains(t, "GO-2", refused[1].Error())
}

func TestObserveNilOnUpHostIsSilent(t *testing.T) {
	r := reacher.New(memory.New())
	assert.Nil(t, r.Observe("sentry", nil))
}
