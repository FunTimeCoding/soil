package job

import "github.com/funtimecoding/soil/pkg/gitlab/constant"

func (j *Job) Fail() bool {
	return j.Status == constant.Failed
}
