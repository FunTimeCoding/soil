package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/failing_jira"
	"testing"
)

func getIssue(
	t *testing.T,
	failure error,
) (*memory.Memory, error) {
	t.Helper()
	c, m := client_tester.NewReported(t, failing_jira.New(failure))
	_, e := c.GetIssue(context.Background(), "ALFA-1", false)

	return m, e
}
