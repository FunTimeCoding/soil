package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestLatestMain(t *testing.T) {
	assert.Any(
		t,
		&pipeline.Pipeline{Reference: "main", Hash: "Bravo"},
		pipeline.LatestMain(
			[]*pipeline.Pipeline{
				{
					Reference: gitConstant.MainBranch,
					Hash:      stringsConstant.UpperAlfa,
				},
				{
					Reference: gitConstant.MainBranch,
					Hash:      stringsConstant.UpperBravo,
				},
				{
					Reference: gitConstant.MainBranch,
					Hash:      stringsConstant.UpperCharlie,
				},
			},
			stringsConstant.UpperBravo,
		),
	)
}
