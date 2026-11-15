package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"go/format"
	"testing"
)

func Formatted(
	t *testing.T,
	content string,
) {
	t.Helper()
	formatted, e := format.Source([]byte(content))
	assert.FatalOnError(t, e)
	assert.String(t, string(formatted), content)
}
