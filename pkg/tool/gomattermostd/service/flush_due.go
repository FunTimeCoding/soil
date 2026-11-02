package service

import "time"

func (s *Service) FlushDue() []error {
	if s.watcher == nil {
		return nil
	}

	var result []error

	for root, events := range s.watcher.Due(time.Now(), s.debounce) {
		result = append(result, s.flush(root, events)...)
	}

	return result
}
