package service

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"

func (s *Service) Pending() ([]*publish.Change, error) {
	result, e := s.store.Unpublished()

	if e != nil {
		return nil, e
	}

	return s.publisher.Changes(result)
}
