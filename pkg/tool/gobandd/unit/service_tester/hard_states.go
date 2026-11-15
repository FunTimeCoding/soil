package service_tester

import bandConstant "github.com/funtimecoding/soil/pkg/band/constant"

func HardStates() []int {
	return []int{bandConstant.PowerCycle, bandConstant.PowerOffSoft}
}
