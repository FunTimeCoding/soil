package base

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/worker"
)

func (s *Server) Refresh(issues ...*issue.Issue) {
	s.t.Helper()
	s.Jira.Issues = issues

	for _, w := range s.Daemon.Workers() {
		if c, okay := w.(*worker.Worker); okay {
			c.Poll()

			return
		}
	}

	s.t.Fatal("the daemon hands out no clock")
}
