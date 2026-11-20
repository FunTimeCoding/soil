package client_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
)

func (o *Tester) GetIssue(failure error) (*memory.Memory, error) {
	o.t.Helper()
	j := mock_client.New(o.t)
	j.Fail(j.Issue, failure)
	c, m := o.Reported(j)
	_, e := c.GetIssue(context.Background(), "ALFA-1", false)

	return m, e
}
