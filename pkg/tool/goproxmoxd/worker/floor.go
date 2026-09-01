package worker

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"

func (w *Worker) Floor() *floor.Floor {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	if w.floor == nil {
		return floor.New()
	}

	return w.floor
}
