package service

import "github.com/funtimecoding/soil/pkg/system/constant"

func launchState(
	identifier string,
	status string,
) string {
	if identifier != constant.LaunchctlAbsent {
		return constant.ServiceRunning
	}

	if status != constant.LaunchctlSuccess {
		return constant.ServiceFailed
	}

	return constant.ServiceStopped
}
