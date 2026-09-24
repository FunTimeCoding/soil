package worker

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/query"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (w *Worker) closedStatus() string {
	return join.Comma(query.Quote(w.client.MustIssueOption().ClosedStatus))
}
