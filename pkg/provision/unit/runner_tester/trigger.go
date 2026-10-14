package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
)

func (o *Tester) Trigger(request trigger.Request) {
	o.t.Helper()
	assert.FatalOnError(o.t, o.Runner.Trigger(request))
}
