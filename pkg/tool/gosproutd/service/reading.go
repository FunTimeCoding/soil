package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/pulse"
)

func (s *Service) Reading(
	session string,
	target *target.Target,
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

	result.Closed = target != nil && target.ClosedAt != nil

	return result
}
