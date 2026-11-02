package terminal

import telemetryConstant "github.com/funtimecoding/soil/pkg/telemetry/constant"

func (t *Terminal) Exit(code int) {
	outcome := telemetryConstant.Success

	if code != 0 {
		outcome = telemetryConstant.Error
	}

	t.end(outcome, code)
}
