package worker

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"

func (w *Worker) Active() bool {
	for _, entry := range w.Entries() {
		if board_entry.Active(entry.Status) {
			return true
		}
	}

	return false
}
