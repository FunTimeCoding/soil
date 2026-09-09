package diff

import "gitlab.com/gitlab-org/api/client-go/v2"

func NewSlice(v []*gitlab.Diff) []*Diff {
	result := make([]*Diff, 0, len(v))

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
