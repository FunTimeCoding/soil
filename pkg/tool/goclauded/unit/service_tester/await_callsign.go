package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (o *Tester) AwaitCallsign(
	x context.Context,
	sessionIdentifier string,
	since time.Time,
) string {
	callsign, e := o.Service.AwaitCallsign(
		x,
		sessionIdentifier,
		since,
		constant.ChannelHold,
	)
	assert.FatalOnError(o.t, e)

	return callsign
}
