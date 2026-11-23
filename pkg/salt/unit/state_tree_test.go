package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/salt/state_tree"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestStateTreeRootsMerge(t *testing.T) {
	assert.Count(
		t,
		0,
		state_tree.New(
			[]string{
				fixture.Path(systemConstant.SaltPath, "secret"),
				fixture.Path(systemConstant.SaltPath, "state"),
			},
			constant.DefaultEnvironments,
		).Analyze().Missing,
	)
}

func TestStateTreeFirstRootWins(t *testing.T) {
	assert.Strings(
		t,
		[]string{"stagingonly"},
		state_tree.New(
			[]string{
				fixture.Path(systemConstant.SaltPath, "secret"),
				fixture.Path(systemConstant.SaltPath, "state"),
			},
			constant.DefaultEnvironments,
		).Analyze().Unreached,
	)
}

func TestStateTreeReached(t *testing.T) {
	assert.Strings(
		t,
		[]string{"common", "common.child", "conditional", "required", "web"},
		state_tree.New(
			[]string{fixture.Path(systemConstant.SaltPath, "state")},
			constant.DefaultEnvironments,
		).Analyze().Reached,
	)
}

func TestStateTreePossible(t *testing.T) {
	assert.Strings(
		t,
		[]string{"helpers.gamma", "service.alpha", "service.beta"},
		state_tree.New(
			[]string{fixture.Path(systemConstant.SaltPath, "state")},
			constant.DefaultEnvironments,
		).Analyze().Possible,
	)
}

func TestStateTreeUnreached(t *testing.T) {
	assert.Strings(
		t,
		[]string{"orphan", "stagingonly"},
		state_tree.New(
			[]string{fixture.Path(systemConstant.SaltPath, "state")},
			constant.DefaultEnvironments,
		).Analyze().Unreached,
	)
}

func TestStateTreeMissing(t *testing.T) {
	assert.Strings(
		t,
		[]string{"absent"},
		state_tree.New(
			[]string{fixture.Path(systemConstant.SaltPath, "state")},
			constant.DefaultEnvironments,
		).Analyze().Missing,
	)
}

func TestStateTreeUnreachedFiles(t *testing.T) {
	assert.Strings(
		t,
		[]string{"unused/draft.txt", "unused/notes.txt"},
		state_tree.New(
			[]string{fixture.Path(systemConstant.SaltPath, "state")},
			constant.DefaultEnvironments,
		).Analyze().UnreachedFiles,
	)
}

func TestStateTreeDeadEnvironments(t *testing.T) {
	assert.Strings(
		t,
		[]string{"staging"},
		state_tree.New(
			[]string{fixture.Path(systemConstant.SaltPath, "state")},
			constant.DefaultEnvironments,
		).Analyze().DeadEnvironments,
	)
}

func TestStateTreeStagingLive(t *testing.T) {
	assert.Strings(
		t,
		[]string{"orphan"},
		state_tree.New(
			[]string{fixture.Path(systemConstant.SaltPath, "state")},
			[]string{"base", "prod", "staging"},
		).Analyze().Unreached,
	)
}
