package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
)

func (o *Tester) Subscribe(callsign string) {
	o.t.Helper()
	_, _, _, e := o.Service().SubscribeMergeRequest(
		o.t.Context(),
		constant.FixtureInstance,
		constant.FixtureProject,
		constant.FixtureRequest,
		callsign,
		"",
		goclaudedConstant.DeliveryImmediate,
	)
	assert.FatalOnError(o.t, e)
}
