package database_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"time"
)

func (o *Tester) replace(
	path string,
	content []byte,
	at time.Time,
) {
	o.t.Helper()
	assert.FatalOnError(o.t, os.WriteFile(path, content, 0o600))
	assert.FatalOnError(o.t, os.Chtimes(path, at, at))
}
