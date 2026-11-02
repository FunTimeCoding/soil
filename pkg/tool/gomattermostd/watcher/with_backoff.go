package watcher

import "time"

func (w *Watcher) WithBackoff(
	minimum time.Duration,
	maximum time.Duration,
) *Watcher {
	w.hold.WithBackoff(minimum, maximum)

	return w
}
