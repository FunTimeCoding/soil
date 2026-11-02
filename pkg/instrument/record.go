package instrument

import (
	telemetryConstant "github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
)

func (i *Instrument) record(
	name string,
	outcome string,
) {
	i.recorder.Record(
		record.NewDomain(
			name,
			telemetryConstant.CommandLine,
			telemetryConstant.User,
			outcome,
		),
	)
}
