package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"time"
)

func (s *Service) flush(
	root string,
	events []*event.Event,
) []error {
	if len(events) == 0 {
		return nil
	}

	subscriptions, e := s.store.ByRoot(root)

	if e != nil {
		return []error{e}
	}

	var result []error

	for _, v := range subscriptions {
		if f := s.notifier.Notify(
			v.Callsign,
			digest.Build(v.Label(), events),
		); f != nil {
			result = append(result, f)
		}
	}

	if f := s.store.TouchRoot(root, time.Now()); f != nil {
		result = append(result, f)
	}

	return result
}
