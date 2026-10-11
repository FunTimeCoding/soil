package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Complete(
	sessionIdentifier string,
	name string,
	topic string,
	message string,
) {
	assert.FatalOnError(
		o.t,
		o.Service.Complete(sessionIdentifier, name, topic, message),
	)
}
