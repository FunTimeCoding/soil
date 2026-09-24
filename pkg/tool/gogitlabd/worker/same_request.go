package worker

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func sameRequest(
	a *merge_request.Request,
	b *merge_request.Request,
) bool {
	return a.Project == b.Project &&
		a.Identifier == b.Identifier &&
		a.Title == b.Title &&
		a.State == b.State
}
