package system

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/minimum"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"slices"
)

func (s *Service) Summaries(m float64) ([]*system_result.Summary, error) {
	if e := minimum.Check(m); e != nil {
		return nil, e
	}

	details, e := s.Details()

	if e != nil {
		return nil, e
	}

	result := make([]*system_result.Summary, 0, len(details))

	for _, d := range details {
		if d.Summary.TopScore >= m {
			result = append(result, d.Summary)
		}
	}

	slices.SortFunc(
		result,
		func(a *system_result.Summary, b *system_result.Summary) int {
			return cmp.Or(
				cmp.Compare(a.Priority, b.Priority),
				cmp.Compare(b.Findings, a.Findings),
				cmp.Compare(a.Name, b.Name),
			)
		},
	)

	return result, nil
}
