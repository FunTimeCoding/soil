package space

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (s *Space) Format(f *option.Format) string {
	return status.New(f).String(s.Name).Format()
}
