package worker

import "github.com/funtimecoding/soil/pkg/gitlab/status"

func (w *Worker) Active() bool {
	for _, entry := range w.Entries() {
		if status.Active(entry.Status) {
			return true
		}
	}

	return false
}
