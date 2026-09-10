package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func lifecycleKinds() []string {
	return []string{
		constant.Register,
		constant.SessionEnd,
		constant.CompleteTimeout,
		constant.InactivityTimeout,
	}
}
