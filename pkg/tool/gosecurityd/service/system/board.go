package system

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/board_result"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"slices"
)

func (s *Service) Board(
	selection *system_selection.Selection,
) (*board_result.Result, error) {
	if e := selection.Check(); e != nil {
		return nil, e
	}

	summaries, e := s.Summaries(0)

	if e != nil {
		return nil, e
	}

	result := board_result.New()

	for _, v := range summaries {
		result.CountFacets(selection, v)

		if !selection.Matches(v) {
			continue
		}

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

		if v.Critical > 0 {
			result.Critical++
		}

		result.Kernels[v.KernelState]++
		result.WithFix += v.Findings - v.NoFix
		result.NoFix += v.NoFix
		result.Rows = append(result.Rows, v)
	}

	slices.SortStableFunc(
		result.Rows,
		func(a *system_result.Summary, b *system_result.Summary) int {
			severe := 0

			if selection.Minimum > 0 {
				severe = cmp.Compare(b.Critical, a.Critical)
			}

			return cmp.Or(
				cmp.Compare(a.Priority, b.Priority),
				severe,
				cmp.Compare(b.Upgradable, a.Upgradable),
				cmp.Compare(
					slices.Index(constant.KernelStateOrder, a.KernelState),
					slices.Index(constant.KernelStateOrder, b.KernelState),
				),
				cmp.Compare(a.Name, b.Name),
			)
		},
	)

	return result, nil
}
