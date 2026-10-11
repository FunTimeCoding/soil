package show

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/systemd/helper"
)

func Parse(m map[string]string) *Result {
	activeState := m["ActiveState"]
	subState := m["SubState"]
	activeEnter := helper.ParseTimestamp(
		m[constant.SystemdActiveEnterTimestamp],
	)
	execMainStart := helper.ParseTimestamp(
		m[constant.SystemdExecMainStartTimestamp],
	)

	return &Result{
		ActiveState:   activeState,
		SubState:      subState,
		ActiveEnter:   activeEnter,
		ExecMainStart: execMainStart,
	}
}
