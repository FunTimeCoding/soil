package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/delivery"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"time"
)

func render(
	entries []*queue.Entry,
	messages ...*message.Message,
) string {
	stored := map[uint]*message.Message{}

	for _, m := range messages {
		stored[m.Identifier] = m
	}

	var values []queue.Entry

	for _, e := range entries {
		values = append(values, *e)
	}

	return delivery.New(stored, time.UTC).Context(values)
}
