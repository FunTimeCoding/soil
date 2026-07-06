package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/runner"
)

func (o *Tester) Sync() *runner.SyncResult {
	o.t.Helper()
	result, e := o.Runner.Sync()
	assert.FatalOnError(o.t, e)

	return result
}
