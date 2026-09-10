package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/finding"
)

func (s *Service) missingTranscriptFindings() ([]*finding.Finding, error) {
	sessions, e := s.store.AllSessions(0, 0)

	if e != nil {
		return nil, e
	}

	var result []*finding.Finding
	settled := s.clock().Add(-constant.CompleteTimeoutWindow)

	for _, i := range sessions {
		if _, found := s.cache.Get(i.Identifier); found {
			continue
		}

		if i.LastSeen.After(settled) {
			continue
		}

		refused, refusalError := s.EmptyRefusal(&i)

		if refusalError != nil {
			return nil, refusalError
		}

		if refused == nil {
			result = append(
				result,
				finding.New(
					constant.MissingTranscriptEmpty,
					i.Identifier,
					"no transcript, empty, the sweep will remove it",
					1,
				),
			)

			continue
		}

		result = append(
			result,
			finding.New(
				constant.MissingTranscriptKept,
				i.Identifier,
				refused.Error(),
				1,
			),
		)
	}

	return result, nil
}
