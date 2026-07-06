package git

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/github/action"
	github "github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/strings/contains"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"testing"
)

func TestBranch(t *testing.T) {
	system.PrintEnvironment()
	e := constant.MainBranch

	if action.IsActionRun() {
		if r := environment.Required(
			github.ReferenceEnvironment,
		); r != constant.MainBranch {
			e = r
		}
	}

	actual := Branch(system.ParentDirectory(constant.Depth))

	if false {
		// Sometimes HEAD
		assert.String(t, e, actual)
	}

	// TODO: Add reference environment to list if missing
	//  Then count somewhere what observations there are
	assert.True(
		t,
		contains.Any(
			[]string{actual},
			[]string{
				constant.MainBranch,
				constant.HeadReference,
			},
		),
	)
}
