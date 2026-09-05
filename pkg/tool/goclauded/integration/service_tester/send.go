package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Send(
	name string,
	to string,
	body string,
) {
	assert.FatalOnError(o.t, o.Service.Send(name, to, body))
}
