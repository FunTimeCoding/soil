package database_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"testing"
	"time"
)

func Copy(
	t *testing.T,
	source string,
	destination string,
	at time.Time,
) {
	t.Helper()
	content, e := os.ReadFile(source)
	assert.FatalOnError(t, e)
	replace(t, destination, content, at)
}
