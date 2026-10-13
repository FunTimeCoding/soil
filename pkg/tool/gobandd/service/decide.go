package service

import (
	"github.com/funtimecoding/soil/pkg/band"
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"slices"
)

func Decide(
	operation string,
	force bool,
	available []int,
) (int, error) {
	state, e := requestedState(operation, force)

	if e != nil {
		return 0, e
	}

	if slices.Contains(available, state) {
		return state, nil
	}

	name := band.PowerStateName(state)

	if state == bandConstant.PowerOffSoftGraceful ||
		state == bandConstant.PowerResetGraceful {
		hard, f := requestedState(operation, true)

		if f != nil {
			return 0, f
		}

		if slices.Contains(available, hard) {
			return 0, conflict.Format(
				"%s not offered: %s",
				name,
				constant.GracefulRefusal,
			)
		}
	}

	return 0, conflict.Format("%s not offered in the current power state", name)
}
