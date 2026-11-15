package runner_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) WaitForFinish() {
	o.t.Helper()
	assert.Signalled(o.t, o.Runner.Finished())
}
