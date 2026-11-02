package unit

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newConfluenceClient(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()
	result := basic.New(
		upstream_tester.Locator(t, server),
		"alfa",
		"bravo-token",
	)
	t.Cleanup(result.Close)

	return result
}
