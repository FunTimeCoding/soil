package unit

import bandConstant "github.com/funtimecoding/soil/pkg/band/constant"

func runningStates() []int {
	return []int{
		bandConstant.PowerCycle,
		bandConstant.PowerOffSoft,
		bandConstant.PowerOffSoftGraceful,
		bandConstant.PowerResetGraceful,
	}
}
