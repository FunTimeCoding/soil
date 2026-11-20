package system_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
)

func (o *Tester) Unreadable(directory string) {
	o.t.Helper()
	aside := join.Empty(directory, ".aside")
	assert.FatalOnError(o.t, os.Rename(directory, aside))
	assert.FatalOnError(o.t, os.WriteFile(directory, nil, 0o644))
	o.t.Cleanup(
		func() {
			assert.FatalOnError(o.t, os.Remove(directory))
			assert.FatalOnError(o.t, os.Rename(aside, directory))
		},
	)
}
