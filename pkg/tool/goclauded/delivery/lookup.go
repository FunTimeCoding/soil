package delivery

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (d *Delivery) lookup(e queue.Entry) *message.Message {
	if e.MessageIdentifier == nil {
		return nil
	}

	return d.messages[*e.MessageIdentifier]
}
