package delivery_tester

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func MessageEntry(m *message.Message) *queue.Entry {
	result := Entry(constant.QueueMessage, join.Empty(m.FromName, ": ", m.Body))
	result.MessageIdentifier = new(m.Identifier)

	return result
}
