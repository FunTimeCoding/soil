package watcher

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"

func (w *Watcher) Drain() map[string][]*event.Event {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	result := w.buffer
	w.buffer = map[string][]*event.Event{}
	clear(w.since)

	return result
}
