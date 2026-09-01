package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/keepass"
)

func (s *Service) Delete(identifier string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.refresh()
	entry, group := s.client.EntryByIdentifier(identifier)

	if entry == nil {
		return not_found.New("entry", identifier)
	}

	keepass.RemoveEntry(group, identifier)

	return s.persist()
}
