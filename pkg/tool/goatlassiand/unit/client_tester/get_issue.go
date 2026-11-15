package client_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"testing"
)

func GetIssue(
	t *testing.T,
	failure error,
) (*memory.Memory, error) {
	t.Helper()
	j := service_tester.NewJira(t)
	j.Fail(j.Issue, failure)
	c, m := NewReported(t, j)
	_, e := c.GetIssue(context.Background(), "ALFA-1", false)

	return m, e
}
