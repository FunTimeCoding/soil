package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/types/summary"
)

func (s *Service) Latest(x context.Context) ([]*summary.Summary, error) {
	stanzas, e := guarded(x, s.guard, s.client.Stanzas)

	if e != nil {
		return nil, e
	}

	var result []*summary.Summary

	for _, t := range stanzas {
		m := summary.New(t.Name, t.Status.Code)

		for _, b := range t.Backup {
			if l, okay := m.Latest[b.Type]; okay &&
				l.Timestamp.Stop >= b.Timestamp.Stop {
				continue
			}

			m.Latest[b.Type] = b
		}

		result = append(result, m)
	}

	return result, nil
}
