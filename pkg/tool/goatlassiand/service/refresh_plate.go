package service

import "slices"

func (s *Service) RefreshPlate() []error {
	p, e := s.Plate()

	if e != nil {
		return []error{e}
	}

	s.mutex.Lock()
	previous := s.current
	s.current = p
	s.mutex.Unlock()

	if !slices.EqualFunc(previous.Issues, p.Issues, sameIssue) ||
		!slices.EqualFunc(previous.WatchedIssues, p.WatchedIssues, sameIssue) ||
		!slices.EqualFunc(previous.Newest, p.Newest, sameIssue) ||
		!slices.EqualFunc(previous.Favorites, p.Favorites, samePage) ||
		!slices.EqualFunc(previous.Watched, p.Watched, samePage) {
		s.notifier.Notify()
	}

	return nil
}
