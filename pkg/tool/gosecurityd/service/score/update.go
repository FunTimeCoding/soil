package score

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/mirror"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"time"
)

func (s *Service) update(
	x context.Context,
	m *mirror.Mirror,
	now time.Time,
) error {
	for from := m.Synced.Add(-constant.ScoreOverlap); from.Before(now); {
		to := from.Add(vulnerabilityConstant.NationalDatabaseWindow)

		if to.After(now) {
			to = now
		}

		for start := 0; ; {
			p, e := s.page(x, start, from, to)

			if e != nil {
				return e
			}

			if f := s.apply(p.Scores, m); f != nil {
				return f
			}

			start += len(p.Scores)

			if len(p.Scores) == 0 || start >= p.Total {
				break
			}
		}

		from = to
	}

	m.Synced = now

	return s.apply(nil, m)
}
