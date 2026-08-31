package worker

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (w *Worker) Watched() []*page.Page {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.watched
}
