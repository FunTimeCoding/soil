package response

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func NewJob(j *job.Job) *Job {
	return &Job{
		Identifier: j.Identifier,
		Name:       j.Name,
		Status:     j.Status,
		Stage:      j.Stage,
		Create:     j.Create,
		Link:       j.Link,
		Trace:      j.Trace,
	}
}
