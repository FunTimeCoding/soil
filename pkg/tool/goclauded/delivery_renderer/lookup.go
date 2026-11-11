package delivery_renderer

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (r *Renderer) lookup(e queue.Entry) *message.Message {
	if e.MessageIdentifier == nil {
		return nil
	}

	return r.messages[*e.MessageIdentifier]
}
