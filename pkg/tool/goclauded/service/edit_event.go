package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"
)

func (s *Service) EditEvent(
	identifier uint,
	message string,
) (*event.Event, error) {
	edited, e := s.store.EditEvent(identifier, message)

	if e != nil {
		return nil, e
	}

	switch edited.Kind {
	case constant.Summarize:
		s.store.UpdateSummaryBody(edited.SessionIdentifier, message)
		slug, metadata, f := s.sessionMetadata(edited.SessionIdentifier)

		if f != nil {
			return edited, f
		}

		if slug != "" {
			s.report(s.pushSummary(slug, message, metadata))
		}
	case constant.Complete:
		topic := edited.Metadata[constant.Topic]
		s.store.UpdateCompletionSummary(
			edited.SessionIdentifier,
			topic,
			message,
		)
		slug, metadata, f := s.sessionMetadata(edited.SessionIdentifier)

		if f != nil {
			return edited, f
		}

		if slug != "" {
			c, f := s.store.CompletionByTopic(edited.SessionIdentifier, topic)

			if f == nil && c != nil {
				s.report(s.pushCompletion(slug, c.Sequence, message, metadata))
			}
		}
	}

	s.notify()

	return edited, nil
}
