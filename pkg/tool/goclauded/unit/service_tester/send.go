package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Send(
	name string,
	to string,
	body string,
) {
	_, e := o.Service.Send(name, to, body, false)
	assert.FatalOnError(o.t, e)
}
