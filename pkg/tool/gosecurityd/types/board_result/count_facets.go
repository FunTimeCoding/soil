package board_result

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
)

func (r *Result) CountFacets(
	s *system_selection.Selection,
	v *system_result.Summary,
) {
	if s.WithCoverage("").Matches(v) {
		r.CoverageCounts[""]++

		for show, coverage := range constant.CoverageFilters {
			if show != "" && coverage == v.Coverage {
				r.CoverageCounts[show]++
			}
		}
	}

	if s.WithMinimum(0).Matches(v) {
		r.ScoreCounts[0]++

		if v.TopScore >= constant.CriticalScore {
			r.ScoreCounts[constant.CriticalScore]++
		}
	}

	for key, value := range v.Labels {
		if !s.WithLabel(key, "").Matches(v) {
			continue
		}

		if r.LabelCounts[key] == nil {
			r.LabelCounts[key] = map[string]int{}
		}

		r.LabelCounts[key][""]++
		r.LabelCounts[key][value]++
	}
}
