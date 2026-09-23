package status

import "github.com/funtimecoding/soil/pkg/gitlab/constant"

func Active(v string) bool {
	switch v {
	case constant.JobRunning,
		constant.JobPending,
		constant.JobCreated,
		constant.JobPreparing,
		constant.JobWaitingForResource:

		return true
	}

	return false
}
