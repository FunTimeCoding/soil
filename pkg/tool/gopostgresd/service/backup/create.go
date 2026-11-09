package backup

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/types/backup"
)

func (s *Service) Create(
	x context.Context,
	instance string,
) (*backup.Backup, error) {
	i, e := s.find(instance)

	if e != nil {
		return nil, e
	}

	at := s.clock().UTC()
	name := at.Format(constant.BackupNameLayout)
	path, f := s.store.Prepare(instance, name)

	if f != nil {
		return nil, f
	}

	y, cancel := context.WithTimeout(x, constant.DumpTimeout)
	defer cancel()
	result, g := s.written(y, i, path)

	if g != nil {
		return nil, errors.Join(g, s.store.Discard(path))
	}

	result.Name = name
	result.Time = at

	return result, nil
}
