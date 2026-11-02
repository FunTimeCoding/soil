package instrument

import telemetryConstant "github.com/funtimecoding/soil/pkg/telemetry/constant"

func (i *Instrument) Flush(v any) {
	if v != nil && i.command != "" {
		i.record(i.command, telemetryConstant.Error)
	}

	i.recorder.Flush()
	i.recorder.Close()
	i.reporter.RecoverFlush(v)
}
