package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/salt/state_tree"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestServedLivePillarOnly(t *testing.T) {
	assert.Strings(
		t,
		[]string{"unused/notes.txt"},
		state_tree.New(
			[]string{fixture.Path(systemConstant.SaltPath, "pillar_root")},
			constant.DefaultEnvironments,
		).Served(),
	)
}

func TestCitedFilesUsed(t *testing.T) {
	s := state_tree.New(
		[]string{fixture.Path(systemConstant.SaltPath, "state")},
		constant.DefaultEnvironments,
	)
	s.Cite([]string{"unused/notes.txt"})
	assert.Strings(t, []string{"unused/draft.txt"}, s.Analyze().UnreachedFiles)
}

func TestCitedPatternUsed(t *testing.T) {
	s := state_tree.New(
		[]string{fixture.Path(systemConstant.SaltPath, "state")},
		constant.DefaultEnvironments,
	)
	s.Cite([]string{"unused/{{ name }}"})
	assert.Count(t, 0, s.Analyze().UnreachedFiles)
}
