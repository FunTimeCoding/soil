package instrument

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/telemetry"
)

type Instrument struct {
	reporter *reporter.Reporter
	recorder *telemetry.Client
}
