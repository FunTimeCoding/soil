package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	telemetryConstant "github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
)

func RecordTelemetry(
	c face.Recorder,
	operation string,
	e error,
) {
	outcome := telemetryConstant.Success

	if e != nil {
		outcome = telemetryConstant.Error
	}

	c.Record(
		record.NewBaseline(
			operation,
			telemetryConstant.WebService,
			"user",
			outcome,
		),
	)
}
