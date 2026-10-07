package service

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/types/change"

func (s *Service) Pending() ([]*change.Change, error) {
	result, e := s.store.Unpublished()

	if e != nil {
		return nil, e
	}

	return s.publisher.Changes(result)
}
