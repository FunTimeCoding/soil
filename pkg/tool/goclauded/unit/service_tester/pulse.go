package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (o *Tester) Pulse(
	sessionIdentifier string,
	d constant.Delivery,
) bool {
	delivered, e := o.Service.SendPulse(sessionIdentifier, "", "news", d)
	assert.FatalOnError(o.t, e)

	return delivered
}
