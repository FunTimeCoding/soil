package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestClientAddsAndRemovesOneValue(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.CheckLive()
	c := s.Client(t)
	assert.FatalOnError(
		t,
		c.SetLabel(a.UUID, constant.FixtureLabelKey, "on_call", "reconciler"),
	)
	assert.FatalOnError(
		t,
		c.AddLabel(
			a.UUID,
			constant.FixtureLabelKey,
			constant.FixtureCompatible,
			"reconciler",
		),
	)
	assert.String(t, "lead,on_call", a.Label(constant.FixtureLabelKey))
	assert.FatalOnError(
		t,
		c.RemoveLabel(a.UUID, constant.FixtureLabelKey, "on_call", "reconciler"),
	)
	assert.String(t, "lead", a.Label(constant.FixtureLabelKey))
}

func TestLabelToolAddsBesideAndReadsCommaForm(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.Label,
		map[string]any{
			constant.Key:   constant.FixtureLabelKey,
			constant.Value: "on_call",
		},
	)
	a.MustCallTool(
		constant.Label,
		map[string]any{
			constant.Key:      constant.FixtureLabelKey,
			constant.LabelAdd: constant.FixtureCompatible,
		},
	)
	assert.String(
		t,
		"role=lead,on_call",
		a.MustCallTool(
			constant.Label,
			map[string]any{constant.Key: constant.FixtureLabelKey},
		),
	)
	a.MustCallTool(
		constant.Label,
		map[string]any{
			constant.Key:         constant.FixtureLabelKey,
			constant.LabelRemove: constant.FixtureCompatible,
		},
	)
	assert.String(t, "on_call", a.Label(constant.FixtureLabelKey))
}
