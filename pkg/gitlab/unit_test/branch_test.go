package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestBranch(t *testing.T) {
	assert.NotNil(
		t,
		branch.New(
			&gitlab.Branch{
				Name:   upper.Alfa,
				Merged: false,
				Commit: &gitlab.Commit{
					CreatedAt: new(constant.StartOfTime),
				},
			},
		),
	)
}
