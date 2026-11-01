package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_list"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	"testing"
)

func renderList(
	t *testing.T,
	l *memory_list.List,
) string {
	t.Helper()
	var b bytes.Buffer
	assert.FatalOnError(t, memory_page.List(l).Render(&b))

	return b.String()
}
