package service

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/outcome"
	"maps"
	"time"
)

func (s *Service) Record(now time.Time) (*outcome.Outcome, error) {
	alerts, e := s.activeAlerts()

	if e != nil {
		return nil, e
	}

	fingerprints, f := s.store.UnresolvedFingerprints()

	if f != nil {
		return nil, f
	}

	open := make(map[string]bool)

	for _, fingerprint := range fingerprints {
		open[fingerprint] = true
	}

	result := outcome.New()
	current := make(map[string]bool)

	for _, a := range alerts {
		if a.State != constant.ActiveState {
			continue
		}

		current[a.Fingerprint] = true

		if open[a.Fingerprint] {
			continue
		}

		if g := s.store.Create(
			record.Record{
				Fingerprint: a.Fingerprint,
				Name:        a.Name,
				Severity:    a.Severity,
				Summary:     a.Summary,
				Labels:      maps.Clone(a.Labels),
				Start:       now,
			},
		); g != nil {
			return nil, g
		}
		open[a.Fingerprint] = true
		result.Created++
	}

	for fingerprint := range open {
		if current[fingerprint] {
			continue
		}

		if h := s.store.Resolve(fingerprint); h != nil {
			return nil, h
		}
	}

	var i error

	if result.Pruned, i = s.store.Prune(now.Add(-s.retention)); i != nil {
		return nil, i
	}

	var j error

	if result.Firing, j = s.store.UnresolvedCount(); j != nil {
		return nil, j
	}

	var k error

	if result.Total, k = s.store.Count(); k != nil {
		return nil, k
	}

	return result, nil
}
