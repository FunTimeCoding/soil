package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) RefusalFor(identifier string) string {
	o.t.Helper()
	r := o.Store.GetSession(identifier)
	assert.True(o.t, r != nil)
	result, e := o.Service.EmptyRefusal(r)
	assert.FatalOnError(o.t, e)
	assert.True(o.t, result != nil)

	return result.Error()
}
