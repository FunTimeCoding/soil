package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Poll() {
	o.t.Helper()
	assert.Count(o.t, 0, o.Service().PollWatches(o.t.Context()))
}
