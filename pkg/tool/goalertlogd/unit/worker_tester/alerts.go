package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
)

func (o *Tester) Alerts(name string) []record.Record {
	o.t.Helper()
	result, e := o.Service.Alerts(name)
	assert.FatalOnError(o.t, e)

	return result
}
