package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestBranch(t *testing.T) {
	assert.NotNil(
		t,
		branch.New(
			&gitlab.Branch{
				Name:   strings.UpperAlfa,
				Merged: false,
				Commit: &gitlab.Commit{CreatedAt: new(constant.StartOfTime)},
			},
		),
	)
}
