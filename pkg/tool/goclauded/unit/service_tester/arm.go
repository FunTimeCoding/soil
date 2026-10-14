package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (o *Tester) Arm(
	sessionIdentifier string,
	mode constant.Cruise,
	pace time.Duration,
) {
	assert.FatalOnError(o.t, o.Service.SetCruise(sessionIdentifier, mode, pace))
}
