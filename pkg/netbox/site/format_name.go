package site

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (s *Site) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", s.Name)
	}

	return s.Name
}
