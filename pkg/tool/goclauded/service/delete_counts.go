package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/receipt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Service) deleteCounts(r *session.Session) (*receipt.Receipt, error) {
	result := receipt.New(r.Identifier, r.Name)

	for _, c := range []countTarget{
		{s.store.CountSessionEvents, &result.Events},
		{s.store.CountSessionEventMetadata, &result.EventMetadata},
		{s.store.CountSessionCompletions, &result.Completions},
		{s.store.CountSessionSummaries, &result.Summaries},
		{s.store.CountSessionLabels, &result.Labels},
		{s.store.CountSessionPulses, &result.Pulses},
		{s.store.CountSessionContextLoads, &result.ContextLoads},
	} {
		count, e := c.read(r.Identifier)

		if e != nil {
			return nil, e
		}

		*c.target = count
	}

	if _, tracked := s.store.TrackerStates()[r.Identifier]; tracked {
		result.TrackerStates = 1
	}

	return result, nil
}
