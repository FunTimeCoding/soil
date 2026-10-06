package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/receipt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/deletion"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Service) deleteCounts(r *session.Session) (*receipt.Receipt, error) {
	result := receipt.New(r.Identifier, r.Name)

	for _, c := range []deletion.CountTarget{
		{Read: s.store.CountSessionEvents, Target: &result.Events},
		{
			Read:   s.store.CountSessionEventMetadata,
			Target: &result.EventMetadata,
		},
		{Read: s.store.CountSessionCompletions, Target: &result.Completions},
		{Read: s.store.CountSessionSummaries, Target: &result.Summaries},
		{Read: s.store.CountSessionLabels, Target: &result.Labels},
		{Read: s.store.CountSessionPulses, Target: &result.Pulses},
		{
			Read:   s.store.CountSessionContextLoads,
			Target: &result.ContextLoads,
		},
	} {
		count, e := c.Read(r.Identifier)

		if e != nil {
			return nil, e
		}

		*c.Target = count
	}

	if _, tracked := s.store.TrackerStates()[r.Identifier]; tracked {
		result.TrackerStates = 1
	}

	return result, nil
}
