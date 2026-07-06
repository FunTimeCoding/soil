package job

import "github.com/funtimecoding/soil/pkg/time"

func (j *Job) formatCompletion() string {
	if j.Raw.Status.CompletionTime != nil {
		return j.Raw.Status.CompletionTime.Format(time.DateMinute)
	}

	return ""
}
