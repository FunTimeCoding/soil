package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"maragu.dev/gomponents"
	"strings"
	"testing"
)

func renderedNode(
	t *testing.T,
	n gomponents.Node,
) string {
	t.Helper()
	var b strings.Builder
	assert.FatalOnError(t, n.Render(&b))

	return b.String()
}
