package mock_client

import (
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
)

func requestedState(
	operation string,
	force bool,
) (int, error) {
	switch operation {
	case constant.OperationStart:
		return bandConstant.PowerOn, nil
	case constant.OperationShutdown:
		if force {
			return bandConstant.PowerOffSoft, nil
		}

		return bandConstant.PowerOffSoftGraceful, nil
	case constant.OperationReboot:
		if force {
			return bandConstant.PowerCycle, nil
		}

		return bandConstant.PowerResetGraceful, nil
	}

	return 0, validation.New("operation unknown: %s", operation)
}
