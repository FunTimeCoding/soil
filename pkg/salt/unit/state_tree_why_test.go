package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/salt/state_tree"
	"github.com/funtimecoding/soil/pkg/salt/types/chain"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestWhyChainPerTopEntry(t *testing.T) {
	r, e := state_tree.New(
		[]string{fixture.Path(systemConstant.SaltPath, "state")},
		constant.DefaultEnvironments,
	).Why("required")
	assert.FatalOnError(t, e)
	assert.String(t, "reached", r.Status)
	assert.Any(
		t,
		[]*chain.Chain{
			{
				Environment: "base",
				Target:      "*",
				Live:        true,
				States:      []string{"common", "common.child", "required"},
			},
			{
				Environment: "prod",
				Target:      "web*",
				Live:        true,
				States:      []string{"web", "required"},
			},
		},
		r.Chains,
	)
}

func TestWhyDeadEntryChain(t *testing.T) {
	r, e := state_tree.New(
		[]string{fixture.Path(systemConstant.SaltPath, "state")},
		constant.DefaultEnvironments,
	).Why("stagingonly")
	assert.FatalOnError(t, e)
	assert.String(t, "unreached", r.Status)
	assert.Any(
		t,
		[]*chain.Chain{
			{
				Environment: "staging",
				Target:      "st*",
				Live:        false,
				States:      []string{"stagingonly"},
			},
		},
		r.Chains,
	)
}

func TestWhyPossible(t *testing.T) {
	r, e := state_tree.New(
		[]string{fixture.Path(systemConstant.SaltPath, "state")},
		constant.DefaultEnvironments,
	).Why("service.beta")
	assert.FatalOnError(t, e)
	assert.String(t, "possible", r.Status)
	assert.Count(t, 0, r.Chains)
}

func TestWhyOrphan(t *testing.T) {
	r, e := state_tree.New(
		[]string{fixture.Path(systemConstant.SaltPath, "state")},
		constant.DefaultEnvironments,
	).Why("orphan")
	assert.FatalOnError(t, e)
	assert.String(t, "unreached", r.Status)
	assert.Count(t, 0, r.Chains)
}

func TestWhyUnknownState(t *testing.T) {
	_, e := state_tree.New(
		[]string{fixture.Path(systemConstant.SaltPath, "state")},
		constant.DefaultEnvironments,
	).Why("nonexistent")
	assert.True(t, not_found.Is(e))
}
