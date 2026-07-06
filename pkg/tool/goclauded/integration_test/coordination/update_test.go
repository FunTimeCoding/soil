package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestUpdate(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "first topic")
	a.CheckLive()
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "second topic",
		},
	)
	r := a.CheckLive()
	updates := clientEntriesByKind(r.Entries, constant.QueueSessionUpdate)
	assert.True(t, len(updates) > 0)
	assert.StringContains(t, "second topic", updates[0].Body)
}

func TestUpdateMultiple(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "step one")
	a.CheckLive()
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "step two",
		},
	)
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "step three",
		},
	)
	r := a.CheckLive()
	updates := clientEntriesByKind(r.Entries, constant.QueueSessionUpdate)
	assert.True(t, len(updates) >= 2)
}

func TestUpdateHistoryEvents(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "initial")
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "milestone one",
		},
	)
	r := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "updated milestone one: completed", r)
}

func TestUpdateBeforeAnnounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	assert.StringContains(
		t,
		"announce first",
		a.MustCallToolError(
			constant.Update,
			map[string]any{
				constant.Message: "completed",
				constant.Topic:   "premature update",
			},
		),
	)
}
