package worker

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"

func (w *Worker) Entries() []*board_entry.Entry {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.entries
}
