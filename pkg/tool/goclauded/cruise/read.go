package cruise

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func Read(
	v *session.Session,
	c *cruise_setting.Setting,
) *Reading {
	result := New()
	result.Mode = c.Mode
	result.Pace = c.Pace

	if c.PromotedAt != nil {
		result.PromotedAt = *c.PromotedAt
	}

	result.Closed = v.ClosedAt != nil
	result.StateKnown = v.LastTurnEndAt != nil

	if result.StateKnown {
		result.IdleSince = *v.LastTurnEndAt
		result.Idle = v.LastPromptAt == nil || v.LastTurnEndAt.After(
			*v.LastPromptAt,
		)
	}

	return result
}
