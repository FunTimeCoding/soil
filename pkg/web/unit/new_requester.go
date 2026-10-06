package unit

import (
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
	"time"
)

func newRequester(
	t *testing.T,
	base string,
) *requester.Requester {
	t.Helper()

	return requester.New(
		upstream_tester.Locator(t, base),
	).WithBackoff(time.Millisecond)
}
