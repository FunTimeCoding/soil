package client_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"testing"
)

func GetIssue(
	t *testing.T,
	failure error,
) (*memory.Memory, error) {
	t.Helper()
	j := mock_client.New(t)
	j.Fail(j.Issue, failure)
	c, m := NewReported(t, j)
	_, e := c.GetIssue(context.Background(), "ALFA-1", false)

	return m, e
}
