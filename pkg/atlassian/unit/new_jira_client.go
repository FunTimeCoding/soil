package unit

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newJiraClient(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()

	return basic.New(upstream_tester.Locator(t, server), "alfa", "bravo-token")
}
