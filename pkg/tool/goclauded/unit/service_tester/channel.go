package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Channel(sessionIdentifier string) int {
	drained, e := o.Service.DrainImmediateQueue(sessionIdentifier, "")
	assert.FatalOnError(o.t, e)

	return len(drained)
}
