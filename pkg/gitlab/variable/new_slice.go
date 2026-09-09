package variable

import "gitlab.com/gitlab-org/api/client-go/v2"

func NewSlice(v []*gitlab.ProjectVariable) []*Variable {
	result := make([]*Variable, 0, len(v))

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
