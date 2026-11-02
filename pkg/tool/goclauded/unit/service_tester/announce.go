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
		o.Roster.Announce(sessionIdentifier, name, topic, files),
	)
}
