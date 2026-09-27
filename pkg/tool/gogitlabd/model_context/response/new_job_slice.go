package response

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func NewJobSlice(v []*job.Job) []*Job {
	result := make([]*Job, len(v))

	for i, j := range v {
		result[i] = NewJob(j)
	}

	return result
}
