package terminal

import "github.com/funtimecoding/soil/pkg/telemetry/constant"

func (t *Terminal) Exit(code int) {
	outcome := constant.Success

	if code != 0 {
		outcome = constant.Error
	}

	t.end(outcome, code)
}
