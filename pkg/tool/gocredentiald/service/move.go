package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/keepass"
)

func (s *Service) Move(
	identifier string,
	groupPath string,
) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.refresh()
	entry, source := s.client.EntryByIdentifier(identifier)

	if entry == nil {
		return not_found.New("entry", identifier)
	}

	target := s.client.GroupByPath(groupPath)

	if target == nil {
		return not_found.New("group", groupPath)
	}

	moved := *entry
	keepass.RemoveEntry(source, identifier)
	s.touch(&moved)
	target.Entries = append(target.Entries, moved)

	return s.persist()
}
