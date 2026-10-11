package fixture

import (
	"testing"

	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
)

func LoadAt(
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
