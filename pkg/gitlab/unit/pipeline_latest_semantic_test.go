package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"testing"
)

func TestLatestSemantic(t *testing.T) {
	assert.Any(
		t,
		&pipeline.Pipeline{Reference: "v1.0.2"},
		pipeline.LatestSemantic(
			[]*pipeline.Pipeline{
				{Reference: "v1.0.0"},
				{Reference: "v1.0.2"},
				{Reference: "v1.0.1"},
			},
		),
	)
}
