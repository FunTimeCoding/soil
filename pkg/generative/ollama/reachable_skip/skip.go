package reachable_skip

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"sync"
	"testing"
	"time"
)

var reachable = sync.OnceValue(
	func() bool {
		x, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		return ollama.NewEnvironment().Reachable(x) == nil
	},
)

func Skip(t *testing.T) {
	t.Helper()

	if !reachable() {
		t.Skipf("embed host unreachable: %s", ollama.NewEnvironment().Locator())
	}
}
