package job

import "github.com/funtimecoding/soil/pkg/gitlab/constant"

func (j *Job) formatUser() string {
	if j.Raw.User != nil {
		return j.Raw.User.Username
	}

	return constant.JobNoUser
}
