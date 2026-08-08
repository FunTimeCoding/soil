package workflow

import "github.com/google/go-github/v90/github"

func NewSlice(v []*github.Workflow) []*Workflow {
	var result []*Workflow

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
