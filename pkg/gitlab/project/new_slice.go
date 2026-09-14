package project

import "gitlab.com/gitlab-org/api/client-go/v3"

func NewSlice(v []*gitlab.Project) []*Project {
	var result []*Project

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
