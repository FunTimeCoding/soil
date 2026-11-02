package service

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (s *Service) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", s.Name)
	}

	return s.Name
}
