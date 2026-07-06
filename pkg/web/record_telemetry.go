package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
)

func RecordTelemetry(
	c face.Recorder,
	operation string,
	e error,
) {
	outcome := constant.Success

	if e != nil {
		outcome = constant.Error
	}

	c.Record(
		record.NewBaseline(
			operation,
			constant.WebService,
			"user",
			outcome,
		),
	)
}
