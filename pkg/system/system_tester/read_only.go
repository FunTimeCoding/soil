package system_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
)

func (o *Tester) ReadOnly(directory string) {
	o.t.Helper()
	assert.FatalOnError(o.t, os.Chmod(directory, 0o500))
	o.t.Cleanup(
		func() {
			assert.FatalOnError(o.t, os.Chmod(directory, 0o700))
		},
	)
}
