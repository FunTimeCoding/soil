package web

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func selectJob(jobs []*job.Job, requested int64) *job.Job {
	if len(jobs) == 0 {
		return nil
	}

	for _, j := range jobs {
		if j.Identifier == requested {
			return j
		}
	}

	for _, j := range jobs {
		if j.Status == constant.JobRunning {
			return j
		}
	}

	for _, j := range jobs {
		if j.Status == constant.JobFail {
			return j
		}
	}

	return jobs[len(jobs)-1]
}
