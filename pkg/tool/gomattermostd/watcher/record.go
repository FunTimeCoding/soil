package watcher

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"time"
)

func (w *Watcher) Record(
	root string,
	e *event.Event,
) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.buffer[root] = append(w.buffer[root], e)

	if _, okay := w.since[root]; !okay {
		w.since[root] = time.Now()
	}
}
