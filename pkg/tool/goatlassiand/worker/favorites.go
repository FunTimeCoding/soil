package worker

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (w *Worker) Favorites() []*page.Page {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.favorites
}
