package database_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"testing"
	"time"
)

func replace(
	t *testing.T,
	path string,
	content []byte,
	at time.Time,
) {
	t.Helper()
	assert.FatalOnError(t, os.WriteFile(path, content, 0o600))
	assert.FatalOnError(t, os.Chtimes(path, at, at))
}
