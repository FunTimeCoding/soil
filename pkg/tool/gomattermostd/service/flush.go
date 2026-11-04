package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"time"
)

func (s *Service) flush(
	root string,
	events []*event.Event,
	now time.Time,
) []error {
	subscriptions, e := s.store.ByRoot(root)

	if e != nil {
		return []error{e}
	}

	identifiers := make([]uint, 0, len(subscriptions))

	for _, v := range subscriptions {
		identifiers = append(identifiers, v.Identifier)
	}

	s.pending.Retain(root, identifiers)
	var result []error

	for _, v := range subscriptions {
		combined := append(s.pending.Take(v.Identifier), events...)

		if len(combined) == 0 {
			continue
		}

		f := s.notifier.Notify(
			v.Callsign,
			digest.Build(v.Label(), combined),
			v.DeliveryMode(),
		)
		crossed, refused := s.reacher.Note(v.Callsign, f)

		if crossed != nil {
			s.logger.Plain("%s", crossed)
		}

		if f == nil {
			s.pending.Done(v.Identifier)

			continue
		}

		s.pending.Keep(v.Identifier, root, combined, now)

		if refused != nil {
			result = append(result, refused)
		}
	}

	if len(events) == 0 {
		return result
	}

	if g := s.store.TouchRoot(root, now); g != nil {
		result = append(result, g)
	}

	return result
}
