package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest"
	"time"
)

func (s *Service) Purge() []error {
	if s.watcher == nil {
		return nil
	}

	pruned, e := s.store.Prune(time.Now().Add(-s.purge))

	if e != nil {
		return []error{e}
	}

	var result []error
	dropped := map[string][]string{}

	for _, v := range pruned {
		dropped[v.Callsign] = append(dropped[v.Callsign], v.Label())

		if f := s.forgetIfUnwatched(v.RootIdentifier); f != nil {
			result = append(result, f)
		}
	}

	for callsign, labels := range dropped {
		if f := s.notifier.Notify(
			callsign,
			digest.Dropped(labels, s.purge),
		); f != nil {
			result = append(result, f)
		}
	}

	if len(pruned) > 0 {
		s.logger.Structured("subscription_purge", "count", len(pruned))
	}

	return result
}
