package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/constant"
)

func (o *Tester) ForceClone() {
	o.t.Helper()
	o.WaitForApply(1)
	o.BreakRemote()

	for range constant.RunnerHealThreshold {
		assert.Error(o.t, o.Sync().Error)
	}
}
