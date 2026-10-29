package service

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"

func (s *Service) AlertsAfter(
	identifier uint,
	limit int,
) ([]record.Record, error) {
	return s.store.After(identifier, limit)
}
