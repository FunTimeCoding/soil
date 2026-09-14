package runner

import "gitlab.com/gitlab-org/api/client-go/v3"

func NewSlice(v []*gitlab.Runner) []*Runner {
	var result []*Runner

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
