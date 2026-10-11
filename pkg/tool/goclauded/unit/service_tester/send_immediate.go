package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) SendImmediate(
	name string,
	to string,
	body string,
) {
	_, e := o.Service.Send(name, to, body, true)
	assert.FatalOnError(o.t, e)
}
