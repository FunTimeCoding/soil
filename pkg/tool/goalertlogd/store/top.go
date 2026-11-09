package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/top"
	"time"
)

func (s *Store) Top(
	n int,
	w *window.Window,
) ([]top.Record, error) {
	var rows []top.Row
	q := s.database.
		Model(record.Stub()).
		Select(
			fmt.Sprintf(
				`name,
				COUNT(*) AS count,
				SUM(CASE WHEN ended_at IS NULL THEN 1 ELSE 0 END)
					AS currently_firing,
				COALESCE(AVG(%s), 0) AS average_seconds,
				MAX(severity) AS severity`,
				averageSecondsExpression(s.database),
			),
		)

	if start, started := w.Start(); started {
		q = q.Where("started_at >= ?", start)
	}

	if end, ended := w.End(); ended {
		q = q.Where("started_at < ?", end)
	}

	q = q.Group(constant.NameColumn).Order("count DESC, name")

	if n > 0 {
		q = q.Limit(n)
	}

	if e := q.Scan(&rows).Error; e != nil {
		return nil, e
	}

	result := make([]top.Record, 0, len(rows))

	for _, r := range rows {
		result = append(
			result,
			top.Record{
				Name:  r.Name,
				Count: r.Count,
				AverageDuration: time.Duration(
					r.AverageSeconds * float64(time.Second),
				),
				CurrentlyFiring: r.CurrentlyFiring,
				Severity:        r.Severity,
			},
		)
	}

	return result, nil
}
