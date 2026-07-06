package pipeline

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestLatestMain(t *testing.T) {
	assert.Any(
		t,
		&gitlab.PipelineInfo{Ref: "main", SHA: "Bravo"},
		LatestMain(
			[]*gitlab.PipelineInfo{
				{Ref: constant.MainBranch, SHA: upper.Alfa},
				{Ref: constant.MainBranch, SHA: upper.Bravo},
				{Ref: constant.MainBranch, SHA: upper.Charlie},
			},
			upper.Bravo,
		),
	)
}
