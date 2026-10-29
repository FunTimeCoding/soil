package worker_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Firing() int {
	o.t.Helper()
	result, e := o.Service.Status()
	assert.FatalOnError(o.t, e)

	return result.Firing
}
