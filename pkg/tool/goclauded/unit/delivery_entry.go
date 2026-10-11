package unit

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func deliveryEntry(
	kind string,
	body string,
) *queue.Entry {
	return queue.New("session-1", "Cedar", kind, body)
}
