package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Release(
	sessionIdentifier string,
	name string,
) {
	assert.FatalOnError(o.t, o.Roster.Release(sessionIdentifier, name))
}
