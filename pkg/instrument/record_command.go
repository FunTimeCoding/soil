package instrument

import "github.com/funtimecoding/soil/pkg/telemetry/constant"

func (i *Instrument) RecordCommand(name string) {
	i.record(name, constant.Success)
	i.command = ""
}
