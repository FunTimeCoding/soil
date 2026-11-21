package system_selection

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/minimum"
)

func (s *Selection) Check() error {
	if _, okay := constant.CoverageFilters[s.Show]; !okay {
		return validation.New("coverage filter unknown: %s", s.Show)
	}

	return minimum.Check(s.Minimum)
}
