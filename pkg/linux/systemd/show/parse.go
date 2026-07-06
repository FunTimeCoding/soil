package show

import (
	"github.com/funtimecoding/soil/pkg/linux/systemd/constant"
	"github.com/funtimecoding/soil/pkg/linux/systemd/helper"
)

func Parse(m map[string]string) *Result {
	activeState := m["ActiveState"]
	subState := m["SubState"]
	activeEnter := helper.ParseTimestamp(m[constant.ActiveEnterTimestamp])
	execMainStart := helper.ParseTimestamp(m[constant.ExecMainStartTimestamp])

	return &Result{
		ActiveState:   activeState,
		SubState:      subState,
		ActiveEnter:   activeEnter,
		ExecMainStart: execMainStart,
	}
}
