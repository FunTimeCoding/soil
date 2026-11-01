package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	"strings"
	"testing"
)

func renderedAction(
	t *testing.T,
	a *action.Action,
) string {
	t.Helper()
	var b strings.Builder
	assert.FatalOnError(t, a.Node().Render(&b))

	return b.String()
}
