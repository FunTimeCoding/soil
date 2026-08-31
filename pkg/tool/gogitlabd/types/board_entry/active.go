package board_entry

import "github.com/funtimecoding/soil/pkg/gitlab/constant"

func Active(status string) bool {
	switch status {
	case constant.JobRunning,
		constant.JobPending,
		constant.JobCreated,
		constant.JobPreparing,
		constant.JobWaitingForResource:

		return true
	}

	return false
}
