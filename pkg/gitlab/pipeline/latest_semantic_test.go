package pipeline

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestLatestSemantic(t *testing.T) {
	assert.Any(
		t,
		&gitlab.PipelineInfo{Ref: "v1.0.2"},
		LatestSemantic(
			[]*gitlab.PipelineInfo{
				{Ref: "v1.0.0"},
				{Ref: "v1.0.2"},
				{Ref: "v1.0.1"},
			},
		),
	)
}
