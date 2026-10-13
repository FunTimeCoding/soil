package worker

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_entry"

func sameRequest(
	a *request_entry.Entry,
	b *request_entry.Entry,
) bool {
	return a.Instance == b.Instance &&
		a.Request.Project == b.Request.Project &&
		a.Request.Identifier == b.Request.Identifier &&
		a.Request.Title == b.Request.Title &&
		a.Request.State == b.Request.State
}
