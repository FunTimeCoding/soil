package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/pulse"
	"time"
)

func (s *Service) Reading(
	session string,
	target *connector.Target,
) *pulse.Reading {
	result := pulse.New()
	unpulsed := s.store.UnpulsedAnswers(session)
	result.Unpulsed = len(unpulsed)

	for _, v := range unpulsed {
		if v.AnsweredAt == nil {
			continue
		}

		if result.OldestAt.IsZero() || v.AnsweredAt.Before(result.OldestAt) {
			result.OldestAt = *v.AnsweredAt
		}

		if v.AnsweredAt.After(result.NewestAt) {
			result.NewestAt = *v.AnsweredAt
		}
	}

	setting := s.store.CruiseSetting(session)
	result.Mode = setting.Mode
	result.Pace = time.Duration(setting.Pace)

	if setting.PromotedAt != nil {
		result.PromotedAt = *setting.PromotedAt
	}

	if target == nil {
		return result
	}

	result.Closed = target.ClosedAt != nil
	result.StateKnown = target.LastTurnEndAt != nil

	if result.StateKnown {
		result.IdleSince = *target.LastTurnEndAt
		result.Idle = target.LastPromptAt == nil || target.LastTurnEndAt.After(
			*target.LastPromptAt,
		)
	}

	return result
}
