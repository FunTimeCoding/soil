package unit

import "github.com/funtimecoding/soil/pkg/band/constant"

func runningStates() []int {
	return []int{
		constant.PowerCycle,
		constant.PowerOffSoft,
		constant.PowerOffSoftGraceful,
		constant.PowerResetGraceful,
	}
}
