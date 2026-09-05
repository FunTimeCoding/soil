package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) SendMessage(
	fromName string,
	toName string,
	body string,
) {
	assert.FatalOnError(o.t, o.Store.SendMessage(fromName, toName, body))
}
