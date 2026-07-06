package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Announce(
	name string,
	topic string,
	files string,
) {
	assert.FatalOnError(o.t, o.Store.Announce(name, topic, files))
}
