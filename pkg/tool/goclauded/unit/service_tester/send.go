package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (o *Tester) Send(
	name string,
	to string,
	body string,
) {
	_, e := o.Service.Send(name, addressed(to), body, constant.DeliveryQueued)
	assert.FatalOnError(o.t, e)
}
