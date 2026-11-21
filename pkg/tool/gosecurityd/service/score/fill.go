package score

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/mirror"
	"time"
)

func (s *Service) fill(
	x context.Context,
	m *mirror.Mirror,
	now time.Time,
) error {
	if m.Start == 0 {
		m.Synced = now
	}

	for x.Err() == nil {
		p, e := s.page(x, m.Start, time.Time{}, time.Time{})

		if e != nil {
			return e
		}

		m.Total = p.Total
		m.Start += len(p.Scores)
		m.Filled = len(p.Scores) == 0 || m.Start >= m.Total

		if f := s.apply(p.Scores, m); f != nil {
			return f
		}

		if m.Filled {
			return nil
		}
	}

	return nil
}
