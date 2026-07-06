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
			if g := s.pushSummary(slug, message, metadata); g != nil {
				return edited, g
			}
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
			c, f := s.store.CompletionByTopic(
				edited.SessionIdentifier,
				topic,
			)

			if f == nil && c != nil {
				if g := s.pushCompletion(
					slug,
					c.Sequence,
					message,
					metadata,
				); g != nil {
					return edited, g
				}
			}
		}
	}

	s.notify()

	return edited, nil
}
