package band

import "github.com/funtimecoding/soil/pkg/band/constant"

func PowerStateName(state int) string {
	switch state {
	case constant.PowerOn:
		return "on"
	case 3, 4:
		return "sleep"
	case constant.PowerCycle:
		return "power cycle"
	case 6:
		return "off hard"
	case 7:
		return "hibernate"
	case constant.PowerOffSoft:
		return "off soft"
	case constant.PowerReset:
		return "reset"
	case 11:
		return "diagnostic interrupt"
	case constant.PowerOffSoftGraceful:
		return "off soft graceful"
	case 13:
		return "off hard graceful"
	case constant.PowerResetGraceful:
		return "reset graceful"
	case 15:
		return "power cycle graceful"
	}

	return "unknown"
}
