package source

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (s *Source) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", s.Name)
	}

	return s.Name
}
