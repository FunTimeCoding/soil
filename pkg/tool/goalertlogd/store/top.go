package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/top"
	"time"
)

func (s *Store) Top(
	n int,
	start time.Time,
	end time.Time,
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
		).
		Where("started_at >= ? AND started_at < ?", start, end).
		Group("name").
		Order("count DESC, name")

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
