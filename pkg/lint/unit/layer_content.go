package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"testing"
)

func layerContent(
	t *testing.T,
	l *layer.Layer,
	path string,
) string {
	t.Helper()
	b, e := l.Read(path)
	assert.FatalOnError(t, e)

	return string(b)
}
