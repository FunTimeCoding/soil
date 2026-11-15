package service_tester

import bandConstant "github.com/funtimecoding/soil/pkg/band/constant"

func RunningStates() []int {
	return []int{
		bandConstant.PowerCycle,
		bandConstant.PowerOffSoft,
		bandConstant.PowerOffSoftGraceful,
		bandConstant.PowerResetGraceful,
	}
}
