package worker

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (w *Worker) Newest() []*issue.Issue {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.newest
}
