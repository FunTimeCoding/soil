package instrument

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
)

func NewCommandLine(i *identity.Tool) *Instrument {
	return &Instrument{
		reporter: reporter.New(i.Name()).Start(),
		recorder: telemetry.NewEnvironment().WithTimeout(
			constant.CommandLineTimeout,
		),
	}
}
