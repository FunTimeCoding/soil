package service

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (s *Service) Update(
	identifier string,
	fields map[string]string,
) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.refresh()
	entry, _ := s.client.EntryByIdentifier(identifier)

	if entry == nil {
		return not_found.New("entry", identifier)
	}

	applyFields(entry, fields)
	s.touch(entry)

	return s.persist()
}
