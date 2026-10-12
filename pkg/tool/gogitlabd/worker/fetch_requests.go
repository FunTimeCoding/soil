package worker

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (w *Worker) fetchRequests() []*merge_request.Request {
	assigned, e := w.service.AssignedMergeRequests(false)
	errors.PanicOnError(e)
	reviewing, f := w.service.ReviewingMergeRequests(false)
	errors.PanicOnError(f)

	return MergeRequests(assigned, reviewing)
}
