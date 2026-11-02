package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/worker"
)

func (o *Tester) Refresh(issues ...*issue.Issue) {
	o.t.Helper()
	o.server.Jira.Issues = issues

	for _, w := range o.server.Daemon.Workers() {
		if c, okay := w.(*worker.Worker); okay {
			c.Poll()

			return
		}
	}

	o.t.Fatal("the daemon hands out no clock")
}
