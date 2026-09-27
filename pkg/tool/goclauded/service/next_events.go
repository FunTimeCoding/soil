package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"

func (s *Service) NextEvents(
	name string,
	kinds []string,
	override *uint,
	limit int,
) ([]event.Event, error) {
	subscription, e := s.store.Subscribe(name, kinds)

	if e != nil {
		return nil, e
	}

	from := subscription.EventIdentifier

	if override != nil {
		from = *override
	}

	events, e := s.store.EventsAfter(from, subscription.KindList(), limit)

	if e != nil {
		return nil, e
	}

	if len(events) == 0 {
		return events, nil
	}

	if e := s.store.SetSubscriptionPosition(
		name,
		events[len(events)-1].Identifier,
	); e != nil {
		return nil, e
	}

	return events, nil
}
