package atlassian_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func Confluence(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()
	result := basic.New(locator.NewOrigin(server), "alfa", "bravo-token")
	t.Cleanup(result.Close)

	return result
}
