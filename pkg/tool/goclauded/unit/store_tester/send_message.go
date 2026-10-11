package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
)

func (o *Tester) SendMessage(
	fromName string,
	toName string,
	body string,
) *message.Message {
	result, e := o.Store.SendMessage(fromName, toName, body)
	assert.FatalOnError(o.t, e)

	return result
}
