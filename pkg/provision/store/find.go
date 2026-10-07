package store

import "github.com/funtimecoding/soil/pkg/provision/model/run"

func (s *Store) Find(identifier uint) (*run.Run, error) {
	var result run.Run

	return &result, s.mapper.Table(s.tableName).First(&result, identifier).Error
}
