package store

import (
	"github.com/funtimecoding/soil/pkg/bolt"
	"github.com/funtimecoding/soil/pkg/notation"
	"go.etcd.io/bbolt"
	"sort"
	"time"
)

func (s *Store) Top(
	n int,
	start time.Time,
	end time.Time,
) ([]TopRecord, error) {
	entries := make(map[string]*topEntry)
	e := s.client.View(
		func(t *bbolt.Tx) error {
			b := s.client.Bucket(t, Bucket)

			if b == nil {
				return nil
			}

			return bolt.For(
				b,
				func(
					_ string,
					v []byte,
				) {
					var r Record
					notation.MustDecodeBytes(v, &r, false)

					if r.Start.Before(start) || !r.Start.Before(end) {
						return
					}

					e, exists := entries[r.Name]

					if !exists {
						e = &topEntry{name: r.Name}
						entries[r.Name] = e
					}

					e.count++
					e.severity = r.Severity

					if r.End != nil {
						e.totalDuration += r.End.Sub(r.Start)
						e.resolvedCount++
					} else {
						e.currentlyFiring++
					}
				},
			)
		},
	)

	if e != nil {
		return nil, e
	}

	result := make([]TopRecord, 0, len(entries))

	for _, e := range entries {
		var avg time.Duration

		if e.resolvedCount > 0 {
			avg = e.totalDuration / time.Duration(e.resolvedCount)
		}

		result = append(
			result,
			TopRecord{
				Name:            e.name,
				Count:           e.count,
				AverageDuration: avg,
				CurrentlyFiring: e.currentlyFiring,
				Severity:        e.severity,
			},
		)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Count > result[j].Count
		},
	)

	if n > 0 && n < len(result) {
		result = result[:n]
	}

	return result, nil
}
