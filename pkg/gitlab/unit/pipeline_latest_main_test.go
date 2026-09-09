package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestLatestMain(t *testing.T) {
	assert.Any(
		t,
		&pipeline.Pipeline{Reference: "main", Hash: "Bravo"},
		pipeline.LatestMain(
			[]*pipeline.Pipeline{
				{Reference: constant.MainBranch, Hash: strings.UpperAlfa},
				{Reference: constant.MainBranch, Hash: strings.UpperBravo},
				{Reference: constant.MainBranch, Hash: strings.UpperCharlie},
			},
			strings.UpperBravo,
		),
	)
}
