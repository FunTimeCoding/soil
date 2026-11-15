package utilization_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	address string,
) *utilization.Client {
	t.Helper()
	result := utilization.New(locator.NewOrigin(address))
	t.Cleanup(result.Close)

	return result
}
