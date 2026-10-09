package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	reacherConstant "github.com/funtimecoding/soil/pkg/reacher/constant"
	"github.com/funtimecoding/soil/pkg/reacher/edge"
	"testing"
	"time"
)

func TestEdgeStringAndContext(t *testing.T) {
	down := edge.Down("sentry", constant.Refused)
	assert.String(t, "sentry down: connection refused", down.String())
	c := down.Context()
	assert.String(t, "down", c[reacherConstant.StateKey].(string))
	assert.String(
		t,
		"connection refused",
		c[reacherConstant.ReasonKey].(string),
	)
	up := edge.Up("sentry", 2*time.Minute)
	assert.String(t, "sentry up after 2m0s", up.String())
	c = up.Context()
	assert.String(t, "up", c[reacherConstant.StateKey].(string))
	assert.String(t, "2m0s", c[reacherConstant.DurationKey].(string))
}
