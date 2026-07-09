package run

import "github.com/google/go-github/v89/github"

func NewSlice(v []*github.WorkflowRun) []*Run {
	var result []*Run

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
