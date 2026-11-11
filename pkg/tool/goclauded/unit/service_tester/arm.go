package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Arm(
	sessionIdentifier string,
	grace string,
) {
	assert.FatalOnError(o.t, o.Cruise.Set(sessionIdentifier, grace))
}
