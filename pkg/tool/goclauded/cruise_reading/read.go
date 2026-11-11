package cruise_reading

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"time"
)

func Read(
	v *session.Session,
	grace time.Duration,
) *Reading {
	result := New()
	result.Grace = grace
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
