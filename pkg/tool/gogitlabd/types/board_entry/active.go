package board_entry

import "github.com/funtimecoding/soil/pkg/gitlab/constant"

func (e *Entry) Active() bool {
	switch e.Status {
	case constant.JobRunning,
		constant.JobPending,
		constant.JobCreated,
		constant.JobPreparing,
		constant.JobWaitingForResource:

		return true
	}

	return false
}
