package job

import timeConstant "github.com/funtimecoding/soil/pkg/time/constant"

func (j *Job) formatCompletion() string {
	if j.Raw.Status.CompletionTime != nil {
		return j.Raw.Status.CompletionTime.Format(timeConstant.DateMinute)
	}

	return ""
}
