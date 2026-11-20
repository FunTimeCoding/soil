package layer_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/layer"
)

func (o *Tester) Content(
	l *layer.Layer,
	path string,
) string {
	o.t.Helper()
	b, e := l.Read(path)
	assert.FatalOnError(o.t, e)

	return string(b)
}
