package service

import "github.com/funtimecoding/soil/pkg/system/constant"

func unitState(active string) string {
	switch active {
	case constant.UnitStateActive:
		return constant.ServiceRunning
	case constant.UnitStateFailed:
		return constant.ServiceFailed
	}

	return constant.ServiceStopped
}
