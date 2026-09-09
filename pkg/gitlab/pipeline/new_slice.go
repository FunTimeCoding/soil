package pipeline

import "gitlab.com/gitlab-org/api/client-go/v2"

func NewSlice(v []*gitlab.PipelineInfo) []*Pipeline {
	result := make([]*Pipeline, 0, len(v))

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
