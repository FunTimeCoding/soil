package overview

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
	"slices"
)

func (s *Service) Overview(show string) (*overview_result.Result, error) {
	flags, e := s.flag.Flags(true)

	if e != nil {
		return nil, e
	}

	result := overview_result.New()

	for i := range flags {
		r := s.row(&flags[i])

		if !r.Flag.Mitigated {
			result.Open++

			if r.Flag.Score >= constant.CriticalScore {
				result.Critical++
			}

			switch r.Fix {
			case constant.FixDebian:
				result.DebianFix++
			case constant.FixLocal:
				result.LocalAction++
			}
		}

		if shown(show, r.Flag.Mitigated) {
			result.Rows = append(result.Rows, r)
		}
	}

	slices.SortStableFunc(
		result.Rows,
		func(a *overview_result.Row, b *overview_result.Row) int {
			return cmp.Or(
				cmp.Compare(b.Flag.Score, a.Flag.Score),
				a.Flag.CreatedAt.Compare(b.Flag.CreatedAt),
			)
		},
	)

	return result, nil
}
