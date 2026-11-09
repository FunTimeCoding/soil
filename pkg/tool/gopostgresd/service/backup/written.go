package backup

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/types/backup"
	"time"
)

func (s *Service) written(
	x context.Context,
	i *instance.Instance,
	path string,
) (*backup.Backup, error) {
	if e := s.dump.Dump(x, i, path); e != nil {
		return nil, e
	}

	sum, size, f := s.store.WriteChecksum(path)

	if f != nil {
		return nil, f
	}

	tables, g := s.dump.Count(x, path)

	if g != nil {
		return nil, g
	}

	result := backup.New(i.Name, "", time.Time{}, size)
	result.Checksum = sum
	result.Tables = tables

	return result, nil
}
