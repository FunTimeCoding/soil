package backup

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/peered"
)

func (s *Service) inPlace(
	x context.Context,
	i *instance.Instance,
	p *peered.Peered,
	path string,
) error {
	if e := connected(x, p); e != nil {
		return e
	}

	if f := s.dump.Restore(x, i, i.Database, path, true); f != nil {
		return f
	}

	s.pool.Forget(i.Name, p)

	return nil
}
