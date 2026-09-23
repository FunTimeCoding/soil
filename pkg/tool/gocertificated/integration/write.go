package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"testing"
)

func write(
	t *testing.T,
	path string,
	content string,
) {
	t.Helper()
	assert.FatalOnError(t, os.WriteFile(path, []byte(content), 0600))
}
