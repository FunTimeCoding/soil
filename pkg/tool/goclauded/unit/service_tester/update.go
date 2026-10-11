package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Update(
	sessionIdentifier string,
	name string,
	topic string,
	message string,
	files string,
) {
	assert.FatalOnError(
		o.t,
		o.Service.Update(sessionIdentifier, name, topic, message, files),
	)
}
