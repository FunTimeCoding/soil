package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"testing"
)

func TestBranch(t *testing.T) {
	assert.NotNil(
		t,
		branch.New(
			&gitlab.Branch{
				Name:   stringsConstant.UpperAlfa,
				Merged: false,
				Commit: &gitlab.Commit{CreatedAt: new(soil.StartOfTime)},
			},
		),
	)
}
