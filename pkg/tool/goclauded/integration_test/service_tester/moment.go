package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Moment(
	sessionIdentifier string,
	name string,
	line string,
) {
	assert.FatalOnError(
		o.t,
		o.Service.Moment(sessionIdentifier, name, line),
	)
}
