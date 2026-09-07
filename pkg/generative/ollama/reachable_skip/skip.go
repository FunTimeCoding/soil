package reachable_skip

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"sync"
	"testing"
)

var reachable = sync.OnceValue(
	func() bool { return ollama.NewEnvironment().Reachable() },
)

func Skip(t *testing.T) {
	t.Helper()

	if !reachable() {
		t.Skipf("embed host unreachable: %s", ollama.NewEnvironment().Locator())
	}
}
