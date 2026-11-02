package service

import (
	"context"
	"slices"
)

func (s *Service) Refresh(x context.Context) []error {
	var failures []error
	changed := false

	for _, i := range s.inventory.Instances {
		entries, e := s.boardOf(x, i.Name)

		if e != nil {
			failures = append(failures, e)

			continue
		}

		requests, f := s.requestsOf(x, i.Name)

		if f != nil {
			failures = append(failures, f)

			continue
		}

		s.mutex.Lock()

		if !slices.EqualFunc(s.entries[i.Name], entries, sameEntry) ||
			!slices.EqualFunc(s.requests[i.Name], requests, sameRequest) {
			changed = true
		}

		s.entries[i.Name] = entries
		s.requests[i.Name] = requests
		s.mutex.Unlock()
	}

	if changed {
		s.notifier.Notify()
	}

	return failures
}
