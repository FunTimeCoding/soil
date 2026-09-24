package worker

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func (w *Worker) Requests() []*merge_request.Request {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.requests
}
