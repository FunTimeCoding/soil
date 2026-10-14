package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/types/update"
)

func (o *Tester) Sync() *update.Result {
	o.t.Helper()
	result, e := o.Runner.Sync()
	assert.FatalOnError(o.t, e)

	return result
}
