package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/peer/constant"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
	"testing"
	"time"
)

func TestEdgeStringAndContext(t *testing.T) {
	down := edge.Down("sentry", errorsConstant.Refused)
	assert.String(t, "sentry down: connection refused", down.String())
	c := down.Context()
	assert.String(t, "down", c[constant.StateKey].(string))
	assert.String(t, "connection refused", c[constant.ReasonKey].(string))
	up := edge.Up("sentry", 2*time.Minute)
	assert.String(t, "sentry up after 2m0s", up.String())
	c = up.Context()
	assert.String(t, "up", c[constant.StateKey].(string))
	assert.String(t, "2m0s", c[constant.DurationKey].(string))
}
