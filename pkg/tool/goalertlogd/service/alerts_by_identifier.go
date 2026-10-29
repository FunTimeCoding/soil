package service

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"

func (s *Service) AlertsByIdentifier(
	identifiers []uint,
) ([]record.Record, error) {
	return s.store.ByIdentifiers(identifiers)
}
