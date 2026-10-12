package service

import (
	band "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
)

func requestedState(
	operation string,
	force bool,
) (int, error) {
	switch operation {
	case constant.OperationStart:
		return band.PowerOn, nil
	case constant.OperationShutdown:
		if force {
			return band.PowerOffSoft, nil
		}

		return band.PowerOffSoftGraceful, nil
	case constant.OperationReboot:
		if force {
			return band.PowerCycle, nil
		}

		return band.PowerResetGraceful, nil
	}

	return 0, validation.New("operation unknown: %s", operation)
}
