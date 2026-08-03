package site

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (s *Site) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", s.Name)
	}

	return s.Name
}
