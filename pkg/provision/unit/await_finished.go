package unit

import (
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"testing"
	"time"
)

func awaitFinished(
	t *testing.T,
	r *runner.Runner,
) {
	t.Helper()

	select {
	case <-r.Finished():
	case <-time.After(time.Second):
		t.Fatal("runner did not finish")
	}
}
