package mock_client

import "github.com/funtimecoding/soil/pkg/band/constant"

func settled(requested int) int {
	switch requested {
	case constant.PowerCycle, constant.PowerReset, constant.PowerResetGraceful:
		return constant.PowerOn
	case constant.PowerOffSoftGraceful:
		return constant.PowerOffSoft
	default:
		return requested
	}
}
