package job

import "github.com/funtimecoding/soil/pkg/github/constant"

func (j *Job) Fail() bool {
	return j.Conclusion == constant.FailureConclusion
}
