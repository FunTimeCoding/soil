package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/runner"
)

func (o *Tester) Trigger(request runner.TriggerRequest) {
	o.t.Helper()
	assert.FatalOnError(o.t, o.Runner.Trigger(request))
}
