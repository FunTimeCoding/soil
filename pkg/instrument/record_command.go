package instrument

import telemetryConstant "github.com/funtimecoding/soil/pkg/telemetry/constant"

func (i *Instrument) RecordCommand(name string) {
	i.record(name, telemetryConstant.Success)
	i.command = ""
}
