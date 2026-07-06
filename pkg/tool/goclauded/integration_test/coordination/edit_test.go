package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestEdit(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "some work")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "hasty summary",
		},
	)
	events := s.Store.Events(event_query.New().SetLimit(1))
	assert.Count(t, 1, events)
	a.MustCallTool(
		constant.EditEvent,
		map[string]any{
			constant.Identifier: float64(events[0].Identifier),
			constant.Message:    "corrected summary with design details",
		},
	)
	assert.StringContains(
		t,
		"corrected summary with design details",
		a.MustCallTool(constant.History, map[string]any{}),
	)
}

func TestEditNegativeID(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "bind identity")
	assert.StringContains(
		t,
		"invalid event identifier",
		a.MustCallToolError(
			constant.EditEvent,
			map[string]any{
				constant.Identifier: float64(-1),
				constant.Message:    "should fail",
			},
		),
	)
}
