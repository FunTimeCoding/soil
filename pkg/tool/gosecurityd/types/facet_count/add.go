package facet_count

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
)

func (c *Count) Add(
	s *system_selection.Selection,
	v *system_result.Summary,
) {
	if s.WithCoverage("").Matches(v) {
		c.Coverage[""]++

		for show, coverage := range constant.CoverageFilters {
			if show != "" && coverage == v.Coverage {
				c.Coverage[show]++
			}
		}
	}

	if s.WithMinimum(0).Matches(v) {
		for _, floor := range constant.ScoreFloors {
			if v.TopScore >= floor {
				c.Score[floor]++
			}
		}
	}

	for key, value := range v.Labels {
		if !s.WithLabel(key, "").Matches(v) {
			continue
		}

		if c.Labels[key] == nil {
			c.Labels[key] = map[string]int{}
		}

		c.Labels[key][""]++
		c.Labels[key][value]++
	}
}
