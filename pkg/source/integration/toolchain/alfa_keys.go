package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"testing"
)

func alfaKeys(
	t *testing.T,
	w *index.Workspace,
	k *kind.Kind,
) (string, string) {
	t.Helper()
	unit, okay := w.UnitKey("example/alfa")
	assert.True(t, okay)
	closure, okay := w.ClosureKey("example/alfa", k)
	assert.True(t, okay)

	return unit, closure
}
