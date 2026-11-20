package system

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/board_result"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"slices"
)

func (s *Service) Board(filter string) (*board_result.Result, error) {
	coverage, okay := constant.CoverageFilters[filter]

	if !okay {
		return nil, validation.New("coverage filter unknown: %s", filter)
	}

	summaries, e := s.Summaries()

	if e != nil {
		return nil, e
	}

	result := board_result.New()

	for _, v := range summaries {
		result.Total++

		switch v.Coverage {
		case constant.CoverageSupported:
			result.Supported++
		case constant.CoverageTracked:
			result.Tracked++
		default:
			result.Untracked++
		}

		if v.Stale {
			result.Stale++
		}

		if v.Upgradable > 0 {
			result.Upgradable++
		}

		if coverage == "" || v.Coverage == coverage {
			result.Rows = append(result.Rows, v)
		}
	}

	slices.SortStableFunc(
		result.Rows,
		func(a *system_result.Summary, b *system_result.Summary) int {
			return cmp.Or(
				cmp.Compare(a.Priority, b.Priority),
				cmp.Compare(b.Upgradable, a.Upgradable),
				cmp.Compare(a.Name, b.Name),
			)
		},
	)

	return result, nil
}
