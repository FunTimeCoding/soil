package transcript

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) noteProfile(readable bool) {
	if !readable {
		if s.profileUnreadable.CompareAndSwap(false, true) {
			s.reporter.CaptureException(errors.New(constant.ProfileUnreadable))
		}

		return
	}

	if s.profileUnreadable.CompareAndSwap(true, false) {
		s.logger.Structured("profile text readable again")
	}
}
