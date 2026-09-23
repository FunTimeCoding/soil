package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (o *Tester) AwaitImmediateQueue(
	x context.Context,
	sessionIdentifier string,
	callsign string,
) []queue.Entry {
	drained, e := o.Service.AwaitImmediateQueue(x, sessionIdentifier, callsign)
	assert.FatalOnError(o.t, e)

	return drained
}
