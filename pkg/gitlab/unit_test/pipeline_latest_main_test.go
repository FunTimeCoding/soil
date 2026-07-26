package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestLatestMain(t *testing.T) {
	assert.Any(
		t,
		&gitlab.PipelineInfo{Ref: "main", SHA: "Bravo"},
		pipeline.LatestMain(
			[]*gitlab.PipelineInfo{
				{Ref: constant.MainBranch, SHA: strings.UpperAlfa},
				{Ref: constant.MainBranch, SHA: strings.UpperBravo},
				{Ref: constant.MainBranch, SHA: strings.UpperCharlie},
			},
			strings.UpperBravo,
		),
	)
}
