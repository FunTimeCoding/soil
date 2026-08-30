package web

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	tool "github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
)

func statusIcon(status string) string {
	switch status {
	case constant.JobSuccess:
		return tool.SuccessIcon
	case constant.JobFail:
		return tool.FailIcon
	case constant.JobRunning:
		return tool.RunningIcon
	case constant.JobCanceled, constant.JobSkipped, constant.JobManual:
		return tool.WarningIcon
	default:
		return tool.PendingIcon
	}
}
