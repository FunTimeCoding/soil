package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"testing"
)

func content(
	t *testing.T,
	path string,
) string {
	t.Helper()
	b, e := os.ReadFile(path)
	assert.FatalOnError(t, e)

	return string(b)
}
