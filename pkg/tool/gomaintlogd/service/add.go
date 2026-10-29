package service

import "github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"

func (s *Service) Add(e *entry.Entry) error {
	if f := check(e); f != nil {
		return f
	}

	if f := s.store.Add(e); f != nil {
		return f
	}

	s.notifier.Notify()

	return nil
}
