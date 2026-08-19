package store

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"

func (s *Store) Create(r record.Record) error {
	return s.database.Create(&r).Error
}
