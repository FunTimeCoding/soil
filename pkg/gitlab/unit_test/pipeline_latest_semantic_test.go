package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestLatestSemantic(t *testing.T) {
	assert.Any(
		t,
		&gitlab.PipelineInfo{Ref: "v1.0.2"},
		pipeline.LatestSemantic(
			[]*gitlab.PipelineInfo{
				{Ref: "v1.0.0"},
				{Ref: "v1.0.2"},
				{Ref: "v1.0.1"},
			},
		),
	)
}
