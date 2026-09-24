package worker

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
)

func (w *Worker) fetchNewest(closed string) []*issue.Issue {
	if !w.HasProject() {
		return nil
	}

	result, e := w.client.SearchLimit(
		constant.NewestLimit,
		NewestQuery(w.project, closed),
	)
	errors.PanicOnError(e)

	return result
}
