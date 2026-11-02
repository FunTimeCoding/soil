package worker

import "github.com/funtimecoding/soil/pkg/gitlab/status"

func (w *Worker) Active() bool {
	for _, entry := range w.entries() {
		if status.Active(entry.Status) {
			return true
		}
	}

	return false
}
