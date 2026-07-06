package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Announce(
	sessionIdentifier string,
	name string,
	topic string,
	files string,
) {
	assert.FatalOnError(
		o.t,
		o.Service.Announce(sessionIdentifier, name, topic, files),
	)
}
