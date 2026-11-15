package requester_tester

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"testing"
	"time"
)

func New(
	t *testing.T,
	base string,
) *requester.Requester {
	t.Helper()
	result := requester.New(locator.NewOrigin(base)).WithBackoff(
		time.Millisecond,
	)
	t.Cleanup(result.Close)

	return result
}
