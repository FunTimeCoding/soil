package index_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/index/kind"
)

func (o *Tester) Keys(
	w *index.Workspace,
	k *kind.Kind,
) (string, string) {
	o.t.Helper()
	unit, okay := w.UnitKey("example/alfa")
	assert.True(o.t, okay)
	closure, okay := w.ClosureKey("example/alfa", k)
	assert.True(o.t, okay)

	return unit, closure
}
