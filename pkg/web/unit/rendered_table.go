package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"strings"
	"testing"
)

func renderedTable(
	t *testing.T,
	a *table.Table,
) string {
	t.Helper()
	var b strings.Builder
	assert.FatalOnError(t, a.Node().Render(&b))

	return b.String()
}
