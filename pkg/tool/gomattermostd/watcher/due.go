package watcher

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"time"
)

func (w *Watcher) Due(
	now time.Time,
	window time.Duration,
) map[string][]*event.Event {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	result := map[string][]*event.Event{}

	for root, at := range w.since {
		if now.Sub(at) < window {
			continue
		}

		result[root] = w.buffer[root]
		delete(w.buffer, root)
		delete(w.since, root)
	}

	return result
}
