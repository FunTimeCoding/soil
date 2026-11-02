package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Moment(
	sessionIdentifier string,
	name string,
	line string,
) {
	assert.FatalOnError(o.t, o.Event.Moment(sessionIdentifier, name, line))
}
