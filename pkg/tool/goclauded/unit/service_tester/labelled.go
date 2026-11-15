package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Labelled(key string) string {
	o.t.Helper()
	value, e := o.Label.GetLabel("session-1", key)
	assert.FatalOnError(o.t, e)

	return value
}
