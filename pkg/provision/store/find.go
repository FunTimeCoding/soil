package store

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/relational"
)

func (s *Store) Find(identifier uint) (*run.Run, error) {
	var result run.Run
	found, e := relational.FindOne(
		s.mapper.Table(s.tableName).Where("id = ?", identifier),
		&result,
	)

	if e != nil {
		return nil, e
	}

	if !found {
		return nil, not_found.New("run", identifier)
	}

	return &result, nil
}
