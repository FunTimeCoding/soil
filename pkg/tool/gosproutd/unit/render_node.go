package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"maragu.dev/gomponents"
	"testing"
)

func renderNode(
	t *testing.T,
	n gomponents.Node,
) string {
	t.Helper()
	var b bytes.Buffer
	assert.FatalOnError(t, n.Render(&b))

	return b.String()
}
