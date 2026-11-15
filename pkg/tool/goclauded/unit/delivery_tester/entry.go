package delivery_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func Entry(
	kind string,
	body string,
) *queue.Entry {
	return queue.New("session-1", "Cedar", kind, body)
}
