package system

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system/matched"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"slices"
)

func (s *Service) findings(v *system.System) (*matched.Matched, error) {
	trackers := s.debian.Trackers()
	scores := s.score.Generation()
	s.mutex.Lock()
	entry := s.cache[v.Name]
	s.mutex.Unlock()

	if entry != nil && entry.Current(v, trackers, scores) {
		return entry, nil
	}

	packages, e := s.store.Packages(v.Identifier)

	if e != nil {
		return nil, e
	}

	result := matched.New(v, trackers, scores)

	if source := s.debian.Covering(v.Release); source != nil {
		if t := source.Tracker(); t != nil {
			result.Findings = matchAll(
				t,
				v.Release,
				packages,
				source.Supported(),
			)
		}
	}

	kernel, running, newest, checked := s.matchKernel(v, packages)
	result.Findings = append(result.Findings, kernel...)
	result.RunningKernel = running
	result.NewestKernel = newest
	result.KernelChecked = checked

	for _, f := range result.Findings {
		if c, okay := s.score.Score(f.Vulnerability); okay {
			f.Score = c.Score
			f.Severity = c.Severity
		}
	}

	slices.SortFunc(
		result.Findings,
		func(a *system_result.Finding, b *system_result.Finding) int {
			return cmp.Or(
				cmp.Compare(a.Source, b.Source),
				cmp.Compare(a.Vulnerability, b.Vulnerability),
			)
		},
	)
	s.mutex.Lock()
	s.cache[v.Name] = result
	s.mutex.Unlock()

	return result, nil
}
