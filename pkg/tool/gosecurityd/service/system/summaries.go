package system

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"slices"
)

func (s *Service) Summaries() ([]*system_result.Summary, error) {
	details, e := s.details()

	if e != nil {
		return nil, e
	}

	result := make([]*system_result.Summary, 0, len(details))

	for _, d := range details {
		result = append(result, d.Summary)
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
