package navigation

import "time"

func (w *Watcher) wait(
	until func() bool,
	d time.Duration,
) bool {
	deadline := time.NewTimer(d)
	defer deadline.Stop()

	for {
		w.mutex.Lock()
		reached := until()
		changed := w.changed
		w.mutex.Unlock()

		if reached {
			return true
		}

		select {
		case <-changed:
		case <-deadline.C:
			return false
		}
	}
}
