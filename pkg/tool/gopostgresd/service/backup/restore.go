package backup

import (
	"context"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
)

func (s *Service) Restore(
	x context.Context,
	instance string,
	name string,
	target string,
) (string, error) {
	i, e := s.find(instance)

	if e != nil {
		return "", e
	}

	path, f := s.store.Existing(instance, name)

	if f != nil {
		return "", f
	}

	if target == "" {
		target = join.Empty(i.Database, constant.RestoreSuffix, name)
	}

	p, g := s.pool.Client(instance)

	if g != nil {
		return "", g
	}

	y, cancel := context.WithTimeout(x, constant.RestoreTimeout)
	defer cancel()

	if target == i.Database {
		return target, s.inPlace(y, i, p, path)
	}

	return target, s.beside(y, i, p, target, path)
}
