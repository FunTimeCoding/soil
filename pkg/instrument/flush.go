package instrument

import "github.com/funtimecoding/soil/pkg/telemetry/constant"

func (i *Instrument) Flush(v any) {
	if v != nil && i.command != "" {
		i.record(i.command, constant.Error)
	}

	i.recorder.Flush()
	i.reporter.RecoverFlush(v)
}
