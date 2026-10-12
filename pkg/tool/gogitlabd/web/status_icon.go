package web

import (
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
)

func statusIcon(status string) string {
	switch status {
	case gitlabConstant.JobSuccess:
		return constant.SuccessIcon
	case gitlabConstant.JobFail:
		return constant.FailIcon
	case gitlabConstant.JobRunning:
		return constant.RunningIcon
	case gitlabConstant.JobCanceled, gitlabConstant.JobSkipped, gitlabConstant.JobManual:
		return constant.WarningIcon
	default:
		return constant.PendingIcon
	}
}
