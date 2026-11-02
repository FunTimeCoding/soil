package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/label_holder"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/label_holder_tester"
	"testing"
)

func TestClaimTakesAnEmptyLabel(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	c := s.Client(t)
	claimed, e := c.ClaimLabel(a.UUID, "role", "watch", "reconciler")
	assert.FatalOnError(t, e)
	assert.True(t, claimed)
	assert.String(t, "watch", a.Label("role"))
}

func TestClaimLeavesADeclaredLabelAlone(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	c := s.Client(t)
	assert.FatalOnError(t, c.SetLabel(a.UUID, "role", "bravo", "operator"))
	claimed, e := c.ClaimLabel(a.UUID, "role", "watch", "reconciler")
	assert.FatalOnError(t, e)
	assert.False(t, claimed)
	assert.String(t, "bravo", a.Label("role"))
}

func TestReleaseLeavesARedeclaredLabelAlone(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	c := s.Client(t)
	_, e := c.ClaimLabel(a.UUID, "role", "watch", "reconciler")
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, c.SetLabel(a.UUID, "role", "none", "operator"))
	released, f := c.ReleaseLabel(a.UUID, "role", "watch", "reconciler")
	assert.FatalOnError(t, f)
	assert.False(t, released)
	assert.String(t, "none", a.Label("role"))
}

func TestReleaseRemovesItsOwnLabel(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	c := s.Client(t)
	_, e := c.ClaimLabel(a.UUID, "role", "watch", "reconciler")
	assert.FatalOnError(t, e)
	released, f := c.ReleaseLabel(a.UUID, "role", "watch", "reconciler")
	assert.FatalOnError(t, f)
	assert.True(t, released)
	assert.String(t, "", a.Label("role"))
}

func TestApplyMovesTheLabelFromAnIdleHolder(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	b := s.NewSession(t)
	b.Announce(b.Name(), "working")
	c := s.Client(t)
	assert.FatalOnError(t, c.SetLabel(b.UUID, "role", "watch", "reconciler"))
	plan := label_holder_tester.Choose(
		label_holder_tester.Active(a.UUID, ""),
		label_holder_tester.Idle(b.UUID, "watch"),
	)
	holder, e := label_holder.Apply(
		c,
		plan,
		label_holder_tester.Rule(),
		"reconciler",
	)
	assert.FatalOnError(t, e)
	assert.String(t, a.UUID, holder)
	assert.String(t, "watch", a.Label("role"))
	assert.String(t, "", b.Label("role"))
}
