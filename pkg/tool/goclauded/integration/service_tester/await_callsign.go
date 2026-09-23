package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"time"
)

func (o *Tester) AwaitCallsign(
	x context.Context,
	sessionIdentifier string,
	since time.Time,
) string {
	callsign, e := o.Service.AwaitCallsign(x, sessionIdentifier, since)
	assert.FatalOnError(o.t, e)

	return callsign
}
