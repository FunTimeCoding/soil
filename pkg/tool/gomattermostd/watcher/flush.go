package watcher

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest"
	"time"
)

func (w *Watcher) Flush(root string) {
	w.mutex.Lock()
	events := w.buffer[root]
	delete(w.buffer, root)
	delete(w.timer, root)
	w.mutex.Unlock()

	if len(events) == 0 {
		return
	}

	subscriptions, e := w.store.ByRoot(root)

	if e != nil {
		w.reporter.CaptureException(e)

		return
	}

	for _, v := range subscriptions {
		w.notifier.Notify(v.Callsign, digest.Build(v.Label(), events))
	}

	if f := w.store.TouchRoot(root, time.Now()); f != nil {
		w.reporter.CaptureException(f)
	}
}
