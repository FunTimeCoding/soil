package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
)

func (o *Tester) ReadMessages(identifiers ...uint) []*message.Message {
	result, e := o.Service.ReadMessages(identifiers)
	assert.FatalOnError(o.t, e)

	return result
}
