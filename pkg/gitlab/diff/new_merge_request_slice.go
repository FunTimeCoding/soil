package diff

import "gitlab.com/gitlab-org/api/client-go/v3"

func NewMergeRequestSlice(v []*gitlab.MergeRequestDiff) []*Diff {
	result := make([]*Diff, 0, len(v))

	for _, e := range v {
		result = append(result, NewMergeRequest(e))
	}

	return result
}
