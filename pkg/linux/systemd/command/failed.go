package command

import "github.com/funtimecoding/soil/pkg/linux/constant"

func Failed() []string {
	return []string{
		constant.SystemdCommand,
		constant.SystemdListUnits,
		constant.SystemdOutput,
		constant.SystemdNotation,
		constant.SystemdState,
		constant.SystemdFailedState,
	}
}
