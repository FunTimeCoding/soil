package system_selection

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"strings"
)

func (s *Selection) Matches(v *system_result.Summary) bool {
	if c := constant.CoverageFilters[s.Show]; c != "" && v.Coverage != c {
		return false
	}

	if v.TopScore < s.Minimum {
		return false
	}

	for key, value := range s.Labels {
		if v.Labels[key] != value {
			return false
		}
	}

	return strings.Contains(strings.ToLower(v.Name), strings.ToLower(s.Search))
}
