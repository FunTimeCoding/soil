package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"testing"
	"time"
)

func copyFile(
	t *testing.T,
	source string,
	destination string,
	at time.Time,
) {
	t.Helper()
	content, e := os.ReadFile(source)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, os.WriteFile(destination, content, 0o600))
	assert.FatalOnError(t, os.Chtimes(destination, at, at))
}
