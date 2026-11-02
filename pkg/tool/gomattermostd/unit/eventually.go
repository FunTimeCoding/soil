package unit

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"testing"
	"time"
)

func eventually(
	t *testing.T,
	condition func() bool,
) {
	t.Helper()
	deadline := time.Now().Add(constant.MattermostSocketWait)

	for time.Now().Before(deadline) {
		if condition() {
			return
		}

		time.Sleep(5 * time.Millisecond)
	}

	t.Fatal("the condition never held")
}
