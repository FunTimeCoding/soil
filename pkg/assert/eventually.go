package assert

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	"testing"
)

func Eventually(
	t *testing.T,
	condition func() bool,
) {
	t.Helper()
	wait_tester.Until(t, condition, soil.FixtureSignalWait)
}
