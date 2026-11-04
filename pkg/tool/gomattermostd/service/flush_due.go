package service

import "time"

func (s *Service) FlushDue() []error {
	if s.watcher == nil {
		return nil
	}

	var result []error
	now := time.Now()
	due := s.watcher.Due(now, s.debounce)

	for root, events := range due {
		result = append(result, s.flush(root, events, now)...)
	}

	for _, root := range s.pending.Roots(now) {
		if _, okay := due[root]; okay {
			continue
		}

		result = append(result, s.flush(root, nil, now)...)
	}

	return result
}
