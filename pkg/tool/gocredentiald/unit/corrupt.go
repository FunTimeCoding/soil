package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"testing"
	"time"
)

func corrupt(
	t *testing.T,
	path string,
	at time.Time,
) {
	t.Helper()
	assert.FatalOnError(t, os.WriteFile(path, []byte("corrupt"), 0o600))
	assert.FatalOnError(t, os.Chtimes(path, at, at))
}
