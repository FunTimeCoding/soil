package transcript

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"strings"
)

func (s *Service) notePersistedProfiles(calls []*tool_call.Call) {
	for _, c := range calls {
		if c.Name != constant.ProfileTool ||
			!strings.HasPrefix(resultText(c.Result), constant.PersistedOutput) ||
			s.clock().Sub(callTime(c)) > constant.PersistedProfileWindow {
			continue
		}

		if s.profilePersisted.CompareAndSwap(false, true) {
			s.reporter.CaptureException(errors.New(constant.ProfilePersisted))
		}
	}
}
