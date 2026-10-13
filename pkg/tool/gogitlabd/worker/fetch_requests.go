package worker

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_entry"

func (w *Worker) fetchRequests(instance string) ([]*request_entry.Entry, error) {
	assigned, e := w.service.AssignedMergeRequests(instance, false)

	if e != nil {
		return nil, e
	}

	reviewing, f := w.service.ReviewingMergeRequests(instance, false)

	if f != nil {
		return nil, f
	}

	var result []*request_entry.Entry

	for _, r := range MergeRequests(assigned, reviewing) {
		result = append(result, request_entry.New(instance, r))
	}

	return result, nil
}
