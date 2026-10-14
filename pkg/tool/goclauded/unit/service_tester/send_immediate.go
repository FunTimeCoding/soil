package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (o *Tester) SendImmediate(
	name string,
	to string,
	body string,
) {
	_, e := o.Service.Send(name, to, body, constant.DeliveryImmediate)
	assert.FatalOnError(o.t, e)
}
