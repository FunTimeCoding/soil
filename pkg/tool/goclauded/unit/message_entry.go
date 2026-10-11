package unit

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func messageEntry(m *message.Message) *queue.Entry {
	result := deliveryEntry(
		constant.QueueMessage,
		join.Empty(m.FromName, ": ", m.Body),
	)
	result.MessageIdentifier = new(m.Identifier)

	return result
}
