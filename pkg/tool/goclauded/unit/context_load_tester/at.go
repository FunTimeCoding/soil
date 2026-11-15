package context_load_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"testing"
)

func At(
	t *testing.T,
	loads []context_load.Load,
	index int,
) *context_load.Load {
	t.Helper()

	if index >= len(loads) {
		t.Fatalf("expected at least %d loads, got %d", index+1, len(loads))
	}

	return &loads[index]
}
