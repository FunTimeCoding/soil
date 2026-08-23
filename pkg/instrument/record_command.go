package instrument

import (
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
)

func (i *Instrument) RecordCommand(name string) {
	i.recorder.Record(
		record.NewDomain(
			name,
			constant.CommandLine,
			constant.User,
			constant.Success,
		),
	)
}
