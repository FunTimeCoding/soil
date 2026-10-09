package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/reacher"
	"syscall"
	"testing"
)

func TestObserveClassifiesTransportFailure(t *testing.T) {
	r := reacher.New()
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
	r := reacher.New()
	assert.Nil(t, r.Observe("sentry", not_found.New("issue", "GO-1")))
	down, _ := r.Down("sentry")
	assert.False(t, down)
	r.Fail("sentry", constant.Refused)
	assert.Nil(t, r.Observe("sentry", not_found.New("issue", "GO-1")))
	down, _ = r.Down("sentry")
	assert.True(t, down)
}

func TestObserveNilOnUpHostIsSilent(t *testing.T) {
	r := reacher.New()
	assert.Nil(t, r.Observe("sentry", nil))
}
