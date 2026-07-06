package session

import (
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/iterm/constant"
)

func (s *Session) formatJob(_ *option.Format) string {
	if s.JobName == "" {
		return constant.NoJob
	}

	return s.JobName
}
