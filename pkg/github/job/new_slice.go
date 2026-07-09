package job

import "github.com/google/go-github/v89/github"

func NewSlice(v []*github.WorkflowJob) []*Job {
	var result []*Job

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
