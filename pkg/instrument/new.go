package instrument

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/telemetry"
)

func New(
	i *identity.Tool,
	version string,
) *Instrument {
	return &Instrument{
		reporter: reporter.New(i.Name(), version).Start(),
		recorder: telemetry.NewEnvironment(),
	}
}
