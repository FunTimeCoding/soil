package database_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"time"
)

func (o *Tester) Copy(
	source string,
	destination string,
	at time.Time,
) {
	o.t.Helper()
	content, e := os.ReadFile(source)
	assert.FatalOnError(o.t, e)
	o.replace(destination, content, at)
}
