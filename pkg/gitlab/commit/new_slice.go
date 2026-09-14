package commit

import "gitlab.com/gitlab-org/api/client-go/v3"

func NewSlice(v []*gitlab.Commit) []*Commit {
	var result []*Commit

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
