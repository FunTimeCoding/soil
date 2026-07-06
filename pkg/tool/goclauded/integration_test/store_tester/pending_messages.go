package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
)

func (o *Tester) PendingMessages(name string) []message.Message {
	result, e := o.Store.PendingMessages(name)
	assert.FatalOnError(o.t, e)

	return result
}
