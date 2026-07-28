package job

import "github.com/funtimecoding/soil/pkg/gitlab/constant"

func (j *Job) formatProject() string {
	if j.Project != nil {
		return j.Project.CombinedName()
	}

	return constant.JobNoProject
}
