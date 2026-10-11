package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"time"
)

func storedMessage(
	identifier uint,
	from string,
	body string,
) *message.Message {
	result := message.New(from, "Cedar", body)
	result.Identifier = identifier
	result.CreatedAt = time.Date(2026, 10, 10, 14, 35, 0, 0, time.UTC)

	return result
}
